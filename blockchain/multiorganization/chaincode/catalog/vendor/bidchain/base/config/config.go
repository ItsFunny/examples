/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:50
# @File : go
# @Description :
# @Attention :
*/
package config

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var (
	globakConfig *InMemoryBlockChainBaseConfiguration
)

func init() {
	globakConfig = NewInMemoryBlockChainBaseConfiguration()
}

type BaseConfigServiceImpl struct {
	// *baseImpl.BaseServiceImpl
}

func NewBaseConfigServiceImpl() *BaseConfigServiceImpl {
	l := new(BaseConfigServiceImpl)
	// l.BaseServiceImpl = baseImpl.NewBaseServiceImpl("config")
	return l
}

func AddKey(ot []ObjectType, kg []KeyGenerator) {
	if len(ot) != len(kg) {
		panic("配置错误,ot的长度和kg的长度不匹配")
	}
	for i := 0; i < len(ot); i++ {
		globakConfig.BlockChainKeyContainer[ot[i]] = kg[i]
	}
}

func AddOverChain(keys []OverChainKey, channelName []string, chaincodeNames []string) {
	if len(keys) != len(chaincodeNames) || len(chaincodeNames) != len(channelName) || len(keys) != len(channelName) {
		panic("配置错误,跨链查询相关的都必须长度一直")
	}
	l := len(keys)
	for i := 0; i < l; i++ {
		globakConfig.BlockChainOverContainer[keys[i]] = OverChainStruct{
			ChannelName:   channelName[i],
			ChaincodeName: chaincodeNames[i],
			Args:          nil,
		}
	}
}

// common 配置类,通过map存储参数等转换
type InMemoryBlockChainBaseConfiguration struct {
	// methods        []MethodName
	// argsParameters []ArgsParameter
	// transactionDescription    []TransBaseDescription
	//
	// keyConstants  []ObjectType
	// keyGenerators []KeyGenerater

	// 多版本密钥管理
	// MultiAESSecrets map[uint64]string

	// 用于生成key,是一个map结构,key是常量,value为生成方式
	BlockChainKeyContainer map[ObjectType]KeyGenerator
	// 用于跨链查询
	BlockChainOverContainer map[OverChainKey]OverChainStruct

	// *BaseConfigServiceImpl
}

func NewInMemoryBlockChainBaseConfiguration() *InMemoryBlockChainBaseConfiguration {
	c := new(InMemoryBlockChainBaseConfiguration)
	c.BlockChainOverContainer=make(map[OverChainKey]OverChainStruct)
	c.BlockChainKeyContainer=make(map[ObjectType]KeyGenerator)
	// argsParameter = argsParameter
	// methods = methods
	// transDesc = transDesc
	// c.keyConstants = keyConstants
	// c.keyGenerators = keyGens
	// c.BaseConfigServiceImpl = NewBaseConfigServiceImpl()

	// argsParameter, methods, transDesc, commonDecrypter, keyConstants, keyGens = nil, nil, nil, nil, nil, nil

	return c
}

func GetKey(stub shim.ChaincodeStubInterface, key ObjectType, req ...interface{}) (string, error) {
	return globakConfig.GetKey(stub, key, req...)
}

func GetOverChain(key OverChainKey, args []byte) OverChainStruct {
	obj, exist := globakConfig.BlockChainOverContainer[key]
	if !exist {
		panic("配置错误,无法找到对应的配置,ock为:" + string(key))
	}
	obj.Args = args
	return obj
}

func (c *InMemoryBlockChainBaseConfiguration) GetKey(stub shim.ChaincodeStubInterface, key ObjectType, req ...interface{}) (string, error) {
	if g, exist := c.BlockChainKeyContainer[key]; !exist {
		return "", errors.New(fmt.Sprintf("key=[%v]的生成key函数不存在", key))
	} else {
		return g(stub, key, req...)
	}
}

//
//
// func Encrypt(data []byte, version uint64) ([]byte, error2.IBaseError) {
// 	if nil == InMemoryBlockChainBaseConfiguration {
// 		if e := Config(); nil != e {
// 			return nil, error2.NewConfigError(e, "初始化config错误")
// 		}
// 	}
// 	return InMemoryBlockChainBaseConfiguration.Encrypt(data, version)
// }
//
// func Decrypt(req []byte, version uint64) ([]byte, error2.IBaseError) {
// 	if nil == InMemoryBlockChainBaseConfiguration {
// 		if e := Config(); nil != e {
// 			return nil, error2.NewConfigError(e, "初始化config错误")
// 		}
// 	}
// 	return InMemoryBlockChainBaseConfiguration.Decrypt(req, version)
// }
//
// func GetKey(stub shim.ChaincodeStubInterface, key ObjectType, args ...interface{}) (string, error2.IBaseError) {
// 	if nil == InMemoryBlockChainBaseConfiguration {
// 		if e := Config(); nil != e {
// 			return "", error2.NewConfigError(e, "初始化config错误")
// 		}
// 	}
// 	return InMemoryBlockChainBaseConfiguration.GetKey(stub, key, args...)
// }
//
// func ValidateArguAndReturn(name MethodName, args []string) (BaseFabricAfterValidModel, error2.IBaseError) {
// 	fmt.Println("开始调用config的参数检查")
// 	utils.DebugPrintDetail("=", "配置", InMemoryBlockChainBaseConfiguration)
// 	if nil == InMemoryBlockChainBaseConfiguration {
// 		if e := Config(); nil != e {
// 			return BaseFabricAfterValidModel{}, error2.NewConfigError(e, "初始化config错误")
// 		}
// 	}
// 	return InMemoryBlockChainBaseConfiguration.CheckAndConvt(name, args)
// }
