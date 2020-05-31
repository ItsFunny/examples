/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 15:48 
# @File : base_service.go
# @Description : 业务service的基类
# @Attention : 
*/
package cc

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"myLibrary/go-library/go/converters"
	"net/http"
	"strconv"
	"examples/blockchain/config/common/base/fabric"
	"examples/blockchain/config/common/base/service/impl"
	"examples/blockchain/config/common/constants"
	error2 "examples/blockchain/config/common/error"
)

type IVlinkChainCodeLogicServiceHelper interface {
	Encrypt(bytes []byte, version base.Version) ([]byte, error2.IVlinkError)
	Decrypt(bytes []byte, version base.Version) ([]byte, error2.IVlinkError)

	GetConcreteKey(stub shim.ChaincodeStubInterface, key base.ObjectType, args ...interface{}) (string, error2.IVlinkError)
}

type VlinkChainCodeBaseLogicServiceImpl struct {
	Stub shim.ChaincodeStubInterface
	// 版本信息
	Version uint64
	// 交易的主旨信息
	BaseTransactionType base.TransBaseTypeV2

	ChainCodeHelper IVlinkChainCodeLogicServiceHelper

	*impl.VlinkBaseServiceImpl
}

func NewVlinkChainCodeBaseLogicServiceImpl(init *impl.VlinkBaseServiceImpl, s shim.ChaincodeStubInterface, version uint64, BaseTransactionType base.TransBaseTypeV2) *VlinkChainCodeBaseLogicServiceImpl {
	b := new(VlinkChainCodeBaseLogicServiceImpl)
	b.VlinkBaseServiceImpl = init
	b.Stub = s
	b.Version = version
	b.BaseTransactionType = BaseTransactionType
	return b
}

// ////////////////////////////////////////////////////
// ///////   业务辅助方法
// ////////////////////////////////////////////////////
// ////////////////////////////////////////////////////
func (b *VlinkChainCodeBaseLogicServiceImpl) CheckExist(objectType base.ObjectType, req ...interface{}) (bool, error2.IVlinkError) {
	key, vlinkError := b.ChainCodeHelper.GetConcreteKey(b.Stub, objectType, req...)
	if nil != vlinkError {
		b.Error("获取ot=[%s]的key失败:%s", objectType, vlinkError.Error())
		return false, error2.ErrorsWithMessage(vlinkError, "获取key失败")
	}

	_, bytes, vlinkError := b.GetByKey(base.Key(key))
	if nil != vlinkError {
		b.Error("链上获取数据失败:%s", vlinkError.Error())
		return false, error2.ErrorsWithMessage(vlinkError, "链上查询数据失败")
	} else if nil == bytes || len(bytes) == 0 {
		return false, nil
	}

	return true, nil
}
func (b *VlinkChainCodeBaseLogicServiceImpl) BuildCompositeKey(k base.ObjectType, args ...interface{}) (base.Key, error2.IVlinkError) {
	b.Debug("开始创建组合键,ot=[%+v],args=[%+v]", k, args)
	s, e := b.ChainCodeHelper.GetConcreteKey(b.Stub, k, args...)
	// s, e := b.Stub.CreateCompositeKey(string(k), attributes)
	if nil != e {
		return "", error2.NewFabricError(e, "创建组合键失败")
	}
	b.Debug("创建的组合键为:[%+v]", s)
	return base.Key(s), nil
}
func (b *VlinkChainCodeBaseLogicServiceImpl) PutByKey(baseReq base.BCPutStateReq, data interface{}) error2.IVlinkError {
	if baseReq.Key == "" {
		return error2.NewArguError(nil, "参数key不可为空")
	}
	b.Debug("开始上传,上传信息为: [%v]  ,上传数据为:{%v}", baseReq, data)
	tranTypeLength := byte(len(b.BaseTransactionType))
	bytes := make([]byte, 0)
	// bytes := converter.BigEndianInt642Bytes(int64(b.BaseTransactionType))
	if baseReq.From != "" {
		decodeBytes := []byte(baseReq.From)
		bytes = append(bytes, decodeBytes...)
	} else {
		decodeBytes := make([]byte, constants.FROM_WALLET_ADDRESS_BYTE_LENGTH)
		bytes = append(bytes, decodeBytes...)
	}

	if baseReq.To != "" {
		decodeBytes := []byte(baseReq.To)
		bytes = append(bytes, decodeBytes...)
	} else {
		decodeBytes := make([]byte, constants.TO_WALLET_ADDRESS_BYTE_LENGTH)
		bytes = append(bytes, decodeBytes...)
	}

	amountBytes := converter.BigEndianFloat64ToByte(float64(baseReq.Token))
	bytes = append(bytes, amountBytes...)

	values := make([]byte, 0)
	switch data.(type) {
	case []byte:
		values = append(values, data.([]byte)...)
	default:
		marshal, e := json.Marshal(data)
		if nil != e {
			return error2.NewJSONSerializeError(e, "序列化上链参数失败")
		}
		values = append(values, marshal...)
	}

	v := b.Version
	bytes = append(bytes, converter.BigEndianInt642Bytes(int64(b.Version))...)
	leftBytes := make([]byte, constants.LEFT_BYTE_LENGTH)
	// 设置基本类型长度
	leftBytes[constants.LEFT_BYTE_TYPE_LENGTH_INDEX] = tranTypeLength
	if baseReq.NeedEncrypt {
		leftBytes[constants.LEFT_BYTE_CRYPT_INDEX] = 1
		b.Debug("key=[%s]的数据需要进行加密,版本号位:[%d]", baseReq.Key, v)
		encrypt, vlinkError := b.ChainCodeHelper.Encrypt(values, base.Version(v))
		if nil != vlinkError {
			b.Error("加密失败:%s", vlinkError.Error())
			return error2.ErrorsWithMessage(vlinkError, "加密失败")
		}
		values = encrypt
		b.Debug("[PutByKey] [加密] 开始上传 {key=%s ; baseType=%d ,fromWalletAddress=%s,toWalletAddress=%s,token=%v,version=%v,value=%s} 至区块链", baseReq.Key, b.BaseTransactionType, baseReq.From, baseReq.To, baseReq.Token, v, string(values))
	} else {
		b.Debug("[PutByKey] [非加密] 开始上传 {key=%s ; baseType=%d ,fromWalletAddress=%s,toWalletAddress=%s,token=%v,version=%v,value=%s} 至区块链", baseReq.Key, b.BaseTransactionType, baseReq.From, baseReq.To, baseReq.Token, v, string(values))
	}
	bytes = append(bytes, leftBytes...)

	// 添加空余字节
	idleBytes := make([]byte, constants.IDLE_BYTE_LENGTH)
	bytes = append(bytes, idleBytes...)

	bytes = append(bytes, values...)

	// 添加基本类型字节
	baseTypeBytes := b.BaseTransactionType.BigEndianConvtBytes()
	bytes = append(bytes, baseTypeBytes...)

	if e := b.putByKey(baseReq.Key, bytes); nil != e {
		b.Error("上传数据失败:%s", e.Error())
		return error2.ErrorsWithMessage(e, "上传数据失败")
	}
	b.Debug("成功上传数据,key={%s},总length为:{%d}", baseReq.Key, len(bytes))

	return nil
}
func (b *VlinkChainCodeBaseLogicServiceImpl) putByKey(key base.Key, bytes []byte) error2.IVlinkError {
	if e := b.Stub.PutState(string(key), bytes); nil != e {
		return error2.NewFabricError(e, "插入数据失败")
	}

	return nil
}
func (b *VlinkChainCodeBaseLogicServiceImpl) GetByKey(key base.Key) (base.BCBaseNodeInfo, []byte, error2.IVlinkError) {
	return b.getByKey(string(key))
}
func (b *VlinkChainCodeBaseLogicServiceImpl) GetDecryptDataByKey(key string) (base.BCBaseNodeInfo, []byte, error2.IVlinkError) {
	info, bytes, vlinkError := b.getByKey(key)
	if nil != vlinkError {
		return info, bytes, vlinkError
	}
	if info.Encrypted {
		b.Debug("该数据被加密,开始解密,版本号为:[%v]", info.Version)
		bytes, vlinkError = b.ChainCodeHelper.Decrypt(bytes, info.Version)
		if nil != vlinkError {
			b.Error("解密失败:%s", vlinkError.Error())
			return info, nil, error2.NewFabricError(vlinkError, "解密失败")
		}
		b.Debug("解密成功")
	}
	return info, bytes, nil
}
func (this *VlinkChainCodeBaseLogicServiceImpl) getByKey(key string) (base.BCBaseNodeInfo, []byte, error2.IVlinkError) {
	var (
		result base.BCBaseNodeInfo
	)

	this.Debug("[GetByKey] 开始往区块链中获取数据 key=%s", key)
	bytes, e := this.Stub.GetState(key)
	if nil != e {
		this.Error("[GetByKey] 从区块链上获取数据 {key=%v} 的时候发生了错误:%s", key, e.Error())
		return result, nil, error2.NewFabricError(e, "获取数据失败")
	}
	if nil != bytes && len(bytes) >= constants.VLINK_COMMON_INDEX_END {
		node, modelWallets := base.GetRegularInfoV2(bytes)
		result = node
		if node.Encrypted {
			this.Debug("该key 对应的值被加密,需要进行解密")
			modelWallets, e = this.ChainCodeHelper.Decrypt(modelWallets, node.Version)
			if nil != e {
				this.Error("解密数据失败:{%s}", e.Error())
				return result, modelWallets, error2.NewCryptError(e, 1, "解密数据失败")
			}
			this.Debug("[GetState] [解密后的数据为] 成功从链上获取{key=%s的信息} ,解析得到的 基本类型为:[%d] from钱包地址为:[%s],to钱包地址为:[%s],交易金额为:[%v],版本为:[%v],模型对象的json为:[%s]", key, node.TxBaseType, node.From, node.To, node.Token, node.Version, string(modelWallets))
		} else {
			this.Debug("[GetState] [非加密] 成功从链上获取{key=%s的信息} ,解析得到的 基本类型为:[%d] from钱包地址为:[%s],to钱包地址为:[%s],交易金额为:[%v],版本为:[%v],模型对象的json为:[%s]", key, node.TxBaseType, node.From, node.To, node.Token, node.Version, string(modelWallets))
		}
		// if node.Version != 0 {
		// 	this.Debug("[GetState] 成功从链上获取{key=%s的信息} ,解析得到的 基本类型为:[%d],类型为:[%d] ,from钱包地址为:[%s],to钱包地址为:[%s],交易金额为:[%v],版本为:[%v],模型对象的json为:[%s]", key, node.TxBaseType, node.From, node.Version, node.To, node.Token, hex.EncodeToString(modelWallets))
		// } else {
		// 	this.Debug("[GetState] 成功从链上获取{key=%s的信息} ,解析得到的 基本类型为:[%d],类型为:[%d] ,from钱包地址为:[%s],to钱包地址为:[%s],交易金额为:[%v],版本为:[%v],模型对象的json为:[%s]", key, node.TxBaseType, node.From, node.Version, node.To, node.Token, modelWallets)
		// }
		return result, modelWallets, nil
	} else {
		this.Debug("[GetState] 链上{key=%s}的数据为空", key)
	}

	return result, bytes, nil
}
func (b *VlinkChainCodeBaseLogicServiceImpl) getKey(stub shim.ChaincodeStubInterface, key base.ObjectType, args ...interface{}) (string, error2.IVlinkError) {
	return b.ChainCodeHelper.GetConcreteKey(stub, key, args...)
}

// 跨链调用
func (b *VlinkChainCodeBaseLogicServiceImpl) InvokeOtherCC(req base.InvokeBaseReq) (base.BaseFabricResp, error2.IVlinkError) {
	var args []string
	args = append(args, string(req.MethodName))

	bytes, e := json.Marshal(req.Data)
	if e != nil {
		return base.BaseFabricResp{}, error2.NewJSONSerializeError(e, fmt.Sprintf("序列化data=[%v]", req.Data))
	}
	args = append(args, string(bytes))

	invokeResp := b.Stub.InvokeChaincode(req.ChaincodeName, [][]byte{[]byte(args[0]), []byte(args[1]), []byte(strconv.Itoa(int(b.Version)))}, req.ChannelName)
	b.Debug("invoke 的返回值为:状态码:{%d},返回值:{%v},msg:{%s}", invokeResp.Status, string(invokeResp.Payload), invokeResp.Message)

	if invokeResp.Status != http.StatusOK {
		b.Error("调用链码=[%s],method=[%v],channel=[%v]失败:%s", req.ChaincodeName, req.MethodName, req.ChannelName, invokeResp.GetMessage())
		return base.BaseFabricResp{}, error2.NewFabricError(nil, "区块链调用失败:"+invokeResp.Message)
	}
	resp, err := HandleResponse(invokeResp.Payload)
	if nil != err {
		b.Error("处理返回值失败:%s", e.Error())
		return base.BaseFabricResp{}, error2.ErrorsWithMessage(err, "处理返回值失败")
	}

	return resp, nil
}

func HandleResponse(bytes []byte) (base.BaseFabricResp, error2.IVlinkError) {
	// bytes := response.Payload
	var resp base.BaseFabricResp

	e := json.Unmarshal(bytes, &resp)
	if nil != e {
		return resp, error2.NewJSONSerializeError(e, "反序列化为 BaseFabricResp 失败")
	}

	return resp, nil
}
