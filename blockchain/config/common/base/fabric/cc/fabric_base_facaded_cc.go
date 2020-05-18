/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 13:05 
# @File : fabric_base_facaded_service.go
# @Description : 门面service的baseCC
# @Attention : 
*/
package cc

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"vlink.com/v2/vlink-common/base/fabric"
	"vlink.com/v2/vlink-common/base/service/impl"
	"vlink.com/v2/vlink-common/constants"
	error2 "vlink.com/v2/vlink-common/error"
)

type ArgeChecker func(name base.MethodName) (interface{}, error2.IVlinkError)

type BaseTypeGetter func(name base.MethodName) (base.TransBaseType, error2.IVlinkError)

type IFacadedHandler interface {
	Handler() base.VlinkPeerResponse
}

type IConcreteFacadedService interface {
	HandleDetail(name base.MethodName, req base.BaseFabricAfterValidModel) (base.IVlinkTxBaseResper, error2.IVlinkError)
	SecurityCheckAndConvt(name base.MethodName, args []string) (base.BaseFabricAfterValidModel, error2.IVlinkError)
	// 获取管道ID,不同的channel 拥有着不同的账本,因此查询交易的时候,key也是不同的
	// 2020-01-05 update  需要修改返回值为[]string,存在一条链码部署在多个channel的可能
	GetChannelID() string
}

type VlinkBaseFabricFacadedCC struct {
	Stub shim.ChaincodeStubInterface
	*impl.VlinkBaseServiceImpl
	ConcreteFacadedService IConcreteFacadedService
}

// 2019-12-17 暂时直接返回这个VlinkPeerResponse,如果后续需要修改,直接修改这里即可
func (b *VlinkBaseFabricFacadedCC) Handler() base.VlinkPeerResponse {
	b.BeforeStart("Handler")
	defer b.AfterEnd()

	funcs, args := b.Stub.GetFunctionAndParameters()
	methodName := base.MethodName(funcs)
	b.Debug("开始参数校验")
	reqModel, vlinkError := b.ConcreteFacadedService.SecurityCheckAndConvt(methodName, args)
	if nil != vlinkError {
		b.Error("校验参数错误:%s", vlinkError.Error())
		return b.ReturnResult(nil, vlinkError)
	}
	b.Debug("结束参数校验")

	b.Debug("开始处理请求详情")
	detail, e := b.ConcreteFacadedService.HandleDetail(methodName, reqModel)
	if nil != e {
		b.Error("执行methodName=[%s]发生错误:%s", methodName, e.Error())
		return b.ReturnResult(nil, e)
	}
	detail.SetTxBaseType(reqModel.BaseTransType)
	detail.SetTxDescription(reqModel.BaseTransDescription)
	detail.SetTransactionID(b.Stub.GetTxID())
	detail.SetChannelID(base.ChannelID(b.ConcreteFacadedService.GetChannelID()))
	b.Debug("结束处理请求详情")

	// b.Debug("判断是否需要对交易记录")
	// // 1. 先判断是否是成功的业务
	// if detail.GetCode()&constants.SUCCESS > 0 {
	// 	b.Debug("业务执行成功,判断是否需要记录交易")
	// 	if reqModel.BaseTransType&base.TX_BASE_NEED_RECORD >= base.TX_BASE_NEED_RECORD {
	// 		// 记录交易
	// 		b.Debug("begin 记录本次交易")
	// 		// 可以交给子类实现,也可以本父类实现,图省时,直接父类统一实现
	//
	// 		b.Debug("end 记录本次交易")
	// 	}
	// } else {
	// 	b.Debug("业务执行失败,失败原因:{%s}", detail.GetMsg())
	// }
	// "from":"12jEeDUaMwkrefwcS6AaBzSsFt3xiATcP3","to":"","token":0,"Data":{"dna":"0a5fb72c-f99f-4897-a05d-336923073749","prvKey":"","coinAddress":"1NzvZtHuuDVnyX56j5HaPWBZzrYDTUvrka"},"Code":1,"Msg":"SUCCESS"}
	// 返回结果
	b.Debug("SUCCESS 成功调用[%s],返回值为:[%v]", methodName, detail)

	return b.ReturnResult(detail, nil)
}

func (b *VlinkBaseFabricFacadedCC) ReturnResult(res interface{}, err error2.IVlinkError) base.VlinkPeerResponse {
	if nil != err {
		return base.Fail(err)
	}

	if nil == res {
		return base.Success(impl.SuccessWithEmptyData())
	}

	switch res.(type) {
	case base.IVlinkTxBaseResper:
		t := res.(base.IVlinkTxBaseResper)
		d := t.GetReturnData()

		var (
			dataBytes []byte
			logBytes  []byte
			e         error
		)

		// 1. 先判断是否是成功的业务
		if t.GetCode()&constants.SUCCESS > 0 {
			b.Debug("业务执行成功,判断是否需要记录交易")
			if d != nil {
				dataBytes, e = json.Marshal(d)
				if nil != e {
					return base.Fail(error2.NewJSONSerializeError(e, "序列化失败"))
				}
			}
			b.Debug("判断是否需要对交易记录")
			logInfos := t.GetTXRecordInfoList()
			logBytes, e = json.Marshal(logInfos)
			if nil != e {
				return base.Fail(error2.NewJSONSerializeError(e, "序列化日志记录失败"))
			}

			// if logInfo.BaseType.Contains(base.TX_BASE_NEED_RECORD) {
			// 	fmt.Println("begin basetype")
			// 	for _, c := range logInfo.BaseType {
			// 		fmt.Println(c)
			// 	}
			// 	fmt.Println("end basetype")
			// 	b.Debug("该提案的tx为:{%s},需要进行记录数据", logInfo.BaseType.String())
			// 	logBytes, e = json.Marshal(logInfo)
			// 	if nil != e {
			// 		return base.Fail(error2.NewJSONSerializeError(e, "序列化失败"))
			// 	}
			// }
		} else {
			b.Debug("业务执行失败,失败原因:{%s}", t.GetMsg())
		}
		return base.SuccessWithDetail(dataBytes, logBytes, t.GetCode(), t.GetMsg())
	default:
		// 必须实现该接口,但是不会遇到,因为都有一个外层封装类
		return base.Fail(error2.NewConfigError(nil, "必须实现IVlinkTxBaseResper接口"))
	}
}

// func (b *VlinkBaseFabricFacadedCC) ConfigArgChecker(methods []MethodName, params []ArgsParameter) {
// 	l := len(methods)
// 	for i := 0; i < l; i++ {
// 		b.AddCheck(methods[i], &params[i])
// 	}
// }

// func (b *VlinkBaseFabricFacadedCC) ConfigLogicDesc(methods []MethodName, descs []constants.TransBaseType) {
// 	l := len(methods)
// 	for i := 0; i < l; i++ {
// 		b.AddLogicDesc(methods[i], descs[i])
// 	}
// }

// func NewVlinkBaseFabricFacadedCC(ac ArgeChecker,b BaseTypeGetter) *VlinkBaseFabricFacadedCC {
// 	c := new(VlinkBaseFabricFacadedCC)
// 	c.Log = log.NewVlinkLog()
// 	// c.ArgumentDecrypt = decrypt
// 	// decrypt.SetParent(c)
// 	// c.CheckerAndDecrypter=ac
// 	// c.BaseTyperGetter=b
//
// 	return c
// }
func NewVlinkBaseFabricFacadedCC(stub shim.ChaincodeStubInterface, ConcreteFacadedService IConcreteFacadedService) *VlinkBaseFabricFacadedCC {
	c := new(VlinkBaseFabricFacadedCC)
	c.VlinkBaseServiceImpl = impl.NewVlinkBaseServiceImpl()
	c.Stub = stub
	c.ConcreteFacadedService = ConcreteFacadedService

	return c
}
