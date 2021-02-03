/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/27 10:41
# @File : contract_impl.go
# @Description :
# @Attention :
*/
package baseImpl

import (
	"bidchain/base/config"
	"bidchain/base/models"
	"bidchain/base/services"
	"bidchain/base/services/wrapper"
	"bidchain/fabric/bserror"
	"bidchain/fabric/chaincode"
	"bidchain/fabric/chaincode/ibidchain_contract"
	"bidchain/fabric/context"
	"bidchain/http_framework/protocol"
	"bidchain/micro/catalog/bo"
	"encoding/json"
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

var _ services.IContractAdapter = &ContractAdapterImpl{}

type ContractAdapterImpl struct {
	mediator *chaincode.BasisContract
	*BaseServiceImpl
}

func (c *ContractAdapterImpl) GetByPrefix(objectType config.ObjectType, handler func(shim.StateQueryIteratorInterface) error, args []string) bserror.IBSError {
	c.BeforeStart("GetByPrefix")
	defer c.AfterEnd()
	// StateQueryIteratorInterface

	key, err := c.GetContext().GetStub().GetStateByPartialCompositeKey(string(objectType), args)
	if nil != err {
		return bserror.NewInternalServerError(err, "前缀查询失败,")
	}
	if e := handler(key); nil != e {
		return bserror.NewInternalServerError(err, "前缀查询解析失败,")
	}
	return nil
}

func NewContractAdapterImpl(mediator *chaincode.BasisContract, reqId string) *ContractAdapterImpl {
	var c = &ContractAdapterImpl{
		mediator:        mediator,
		BaseServiceImpl: NewBaseServiceImpl(reqId),
	}
	return c
}

func (c *ContractAdapterImpl) PutInterfaceData(ot config.ObjectType, data interface{}, keyArgs ...interface{}) bserror.IBSError {
	c.BeforeStart("ProtobufStoreInterfaceData")
	defer c.AfterEnd()

	key, ibsError := c.GenCompositeKey(ot, keyArgs...)
	if nil != ibsError {
		c.Error("创建组合键失败:" + ibsError.Error())
		return ibsError
	}

	return c.PutByKey(key, data)
}

func (c *ContractAdapterImpl) PutByKey(key string, data interface{}) bserror.IBSError {
	c.BeforeStart("PutByKey")
	defer c.AfterEnd()

	var e error
	var dataBytes []byte
	switch data.(type) {
	case []byte:
		dataBytes = data.([]byte)
	case proto.Message:
		message := data.(proto.Message)
		data, err := proto.Marshal(message)
		if err != nil {
			e = errors.New("proto序列化失败:" + err.Error())
		} else {
			dataBytes = data
		}
	case string:
		dataBytes = []byte(data.(string))
	default:
		marshal, err := json.Marshal(data)
		if nil != err {
			e = errors.New("json序列化失败:" + err.Error())
		} else {
			dataBytes = marshal
		}
	}
	if nil != e {
		return bserror.NewInternalServerError(e, "上传数据失败,key="+key)
	}
	c.Info("上传数据,key为 [ " + key + " ] ,内容为: [ " + string(dataBytes) + " ]")
	if e = c.GetContext().GetStub().PutState(key, dataBytes); nil != e {
		return bserror.NewInternalServerError(e, "上传数据失败,key="+key)
	}

	return nil
}

func (c *ContractAdapterImpl) GetByBuildKey(ot config.ObjectType, handler func([]byte) error, keyArgs ...interface{}) (string, bserror.IBSError) {
	c.BeforeStart("GetByKey")
	defer c.AfterEnd()
	key, ibsError := c.GenCompositeKey(ot, keyArgs...)
	if nil != ibsError {
		return "", ibsError
	}

	return key, c.GetByKey(key, handler)
}

func (c *ContractAdapterImpl) GetByKey(key string, handler func([]byte) error) bserror.IBSError {
	c.BeforeStart("GetByKey")
	c.AfterEnd()

	state, err := c.GetContext().GetStub().GetState(key)
	if nil != err {
		return bserror.NewInternalServerError(err, "通过key查询数据失败")
	}
	c.Info("查询数据,key:[ "+key+" ],结果为:[ "+string(state)+"] ")
	e := handler(state)
	if nil != e {
		return bserror.NewInternalServerError(e, "error")
	}
	return nil
}

func (c *ContractAdapterImpl) CheckExist(ot config.ObjectType, keyArgs ...interface{}) (wrapper.KeyWrapper, bool, bserror.IBSError) {
	c.BeforeStart("CheckExist")
	defer c.AfterEnd()
	key, ibsError := c.GenCompositeKey(ot, keyArgs...)
	if nil != ibsError {
		return wrapper.KeyWrapper{}, false, ibsError
	}
	c.Info("生成的key为:"+key)
	state, err := c.GetContext().GetStub().GetState(key)
	if nil != err {
		return wrapper.KeyWrapper{}, false, bserror.NewInternalServerError(bserror.INTERNAL_SERVER_ERROR, "链上查询数据失败:"+err.Error())
	}

	return wrapper.KeyWrapper{
		Key: key,
	}, len(state) > 0, nil
}

func (c *ContractAdapterImpl) GetOverChainBaseInfo(key config.OverChainKey, args []byte) config.OverChainStruct {
	c.BeforeStart("GetOverChainBaseInfo")
	defer c.AfterEnd()
	overChain := config.GetOverChain(key, args)

	return overChain
}

func (c *ContractAdapterImpl) CallOverChainWithArgument(key config.OverChainKey, args []byte) ([]byte, bserror.IBSError) {
	c.BeforeStart("CallOverChainWithArgument")
	defer c.AfterEnd()
	res := bo.CatalogAuthenticationResp{
		Allowed: true,
	}
	bs, _ := json.Marshal(res)
	return bs, nil
	//
	// info := c.GetOverChainBaseInfo(key, args)
	// resp, err := c.InvokeOtherChain(models.InvokeOtherChainReq{
	// 	ChannelName:   info.ChannelName,
	// 	ChaincodeName: info.ChaincodeName,
	// 	MethodName:    info.MethodName,
	// 	Args:          info.Args,
	// })
	// if nil != err {
	// 	c.Error("跨链查询失败:" + err.Error())
	// 	return nil, err
	// }
	//
	// return resp.Payload, nil
}

func (c *ContractAdapterImpl) GetContext() context.IBidchainContext {
	return c.mediator.GetContext()
}

func (c *ContractAdapterImpl) GenCompositeKey(objectType config.ObjectType, args ...interface{}) (string, bserror.IBSError) {
	key, err := config.GetKey(c.mediator.GetContext().GetStub(), objectType, args...)
	if nil != err {
		return "", bserror.NewInternalServerError(err, "获取组合键失败")
	}
	c.Info("生成的组合键为:"+key)
	return key, nil
}

func (c *ContractAdapterImpl) InvokeOtherChain(req models.InvokeOtherChainReq) (models.InvokeOtherChainResp, bserror.IBSError) {
	c.BeforeStart("InvokeOtherChain")
	defer c.AfterEnd()

	var (
		res models.InvokeOtherChainResp
	)
	if len(req.ChannelName) == 0 || len(req.ChaincodeName) == 0 {
		c.Error("参数错误,channel名称或者chaincCode名称为空,数据信息为:" + req.String())
		return res, bserror.ARGUMENT_ERROR
	}
	stub := c.mediator.GetContext().GetStub()
	invokeResp := stub.InvokeChaincode(req.ChaincodeName, [][]byte{[]byte(req.MethodName), req.Args}, req.ChannelName)

	c.Info("跨链调用结果为,响应码:" + strconv.Itoa(int(invokeResp.Status)) + ",msg:" + invokeResp.Message + ",返回数据为:" + string(invokeResp.Payload))
	if invokeResp.Status != shim.OK {
		return res, bserror.INTERNAL_SERVER_INVOKE_OTHERCHAIN_ERROR
	}
	res.Payload = invokeResp.Payload

	return res, nil
}

func (c *ContractAdapterImpl) InitBidchainContract(ctx context.IBidchainContext) pb.Response {
	return c.mediator.InitBidchainContract(ctx)
}

func (c *ContractAdapterImpl) InvokeBidchainContract(ctx context.IBidchainContext) pb.Response {
	return c.mediator.InvokeBidchainContract(ctx)
}

func (c *ContractAdapterImpl) GetFunctionByName(name string) string {
	return c.mediator.GetFunctionByName(name)
}

func (c *ContractAdapterImpl) GetChild() ibidchain_contract.IBidchainContract {
	return c.mediator.GetChild()
}

func (c *ContractAdapterImpl) SetChild(contract ibidchain_contract.IBidchainContract) {
	c.mediator.SetChild(contract)
}

func (c *ContractAdapterImpl) GetChaincodeName() string {
	return c.mediator.GetChaincodeName()
}

func (c *ContractAdapterImpl) CommandArrived(request, response protocol.ICommand) {
	c.mediator.CommandArrived(request, response)
}
