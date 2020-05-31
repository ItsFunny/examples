/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-15 11:46 
# @File : base.go
# @Description : 
# @Attention : 
*/
package models

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"myLibrary/go-library/go/crypt"
	"myLibrary/go-library/go/utils"
	"strconv"
	"examples/blockchain/config/common/base/fabric"
	"examples/blockchain/config/common/base/service"
	error2 "examples/blockchain/config/common/error"
)

var (
	COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType base.ObjectType, param ...interface{}) (string, error2.IVlinkError) {
		switch param[0].(type) {
		case string:
			return COMMON_STRING_KEY_GENERATOR(stub, objectType, param[0])
		case []string:
			return COMMON_STRING_ARRAY_KEY_GENERATOR(stub, objectType, param[0])
		}
		return "",error2.ErrorsWithMessage(nil,"找不到匹配的处理")
	}
	COMMON_STRING_ARRAY_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType base.ObjectType, param ...interface{}) (string, error2.IVlinkError) {
		strings := param[0].([]string)
		s, e := stub.CreateCompositeKey(string(objectType), strings)
		if nil != e {
			return "", error2.NewFabricError(e, "创建组合键失败")
		}
		return s, nil
	}

	COMMON_STRING_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType base.ObjectType, param ...interface{}) (string, error2.IVlinkError) {
		strings := param[0].(string)
		s, e := stub.CreateCompositeKey(string(objectType), []string{strings})
		if nil != e {
			return "", error2.NewFabricError(e, "创建组合键失败")
		}
		return s, nil
	}
)

// common 配置类,通过map存储参数等转换
type VlinkCommonBlockchainBaseConfiguration struct {
	methods        []base.MethodName
	argsParameters []ArgsParameter
	baseTypes      []base.TransBaseDescription

	keyConstants  []base.ObjectType
	keyGenerators []base.KeyGenerater

	ArgsDecrypter       ArgsDecrypter
	ArgsCheckMap        map[base.MethodName]*ArgsParameter
	LogicDescriptionMap map[base.MethodName]base.TransBaseDescription

	// 多版本密钥管理
	MultiAESSecrets map[uint64]string

	// 用于生成key,是一个map结构,key是常量,value为生成方式
	BlockChainKeyContainer map[base.ObjectType]base.KeyGenerater
}

func (b *VlinkCommonBlockchainBaseConfiguration) CheckAndConvt(method base.MethodName, args []string) (base.BaseFabricAfterValidModel, error2.IVlinkError) {
	var (
		result base.BaseFabricAfterValidModel
	)

	fmt.Println("检查参数是否安全")
	for _, a := range args {
		fmt.Println(a)
	}

	if p, exist := b.ArgsCheckMap[method]; !exist {
		return result, error2.NewConfigError(nil, "配置错误")
	} else {
		if err := p.ArgsChecker(args); nil != err {
			return result, error2.ErrorsWithMessage(err, "参数checker无法通过")
		}
		if res, vlinkError := p.ArgsConverter(args); nil != vlinkError {
			return result, error2.ErrorsWithMessage(vlinkError, "参数转换无法通过")
		} else {
			// 判断参数是否实现了某个接口
			utils.DebugPrintDetail("+", "请求参数", res)
			switch res.(type) {
			case service.IVlinkValidater:
				if e := res.(service.IVlinkValidater).Validate(); nil != e {
					return result, error2.NewArguError(e, "参数校验错误")
				}
			default:
				return result, error2.NewArguError(nil, "参数必须实现VlinkValidater接口")
			}
			// 参数判断是否需要转换
			switch res.(type) {
			case service.IVlinkCrypter:
				//  0 为参数 1 为version
				d, e := res.(service.IVlinkCrypter).Decrypt(args[1])
				if nil != e {
					return result, error2.NewArguError(e, "参数转换失败")
				}
				result.Req = d
			default:
				result.Req = res
			}
			// 获取baseType
			if baseType, exist := b.LogicDescriptionMap[method]; !exist {
				return result, error2.NewConfigError(nil, "配置错误,baseType未配置")
			} else {
				result.BaseTransType = baseType.TransBaseType
				result.BaseTransDescription = baseType.Description
			}
			v, err := strconv.ParseInt(args[1], 10, 64)
			if nil != err {
				return result, error2.NewArguError(nil, "参数错误,版本号错误:"+args[1])
			}
			result.Version = uint64(v)

			return result, nil
		}
	}

}
func (b *VlinkCommonBlockchainBaseConfiguration) Encrypt(valueBytes []byte, version uint64) ([]byte, error2.IVlinkError) {
	// TODO
	if key, exist := b.MultiAESSecrets[version]; !exist {
		return nil, error2.NewConfigError(nil, "配置错误,无法找到匹配的密钥")
	} else {
		bytes, e := encrypt.AesEncrypt(valueBytes, []byte(key))
		if nil != e {
			return nil, error2.NewArguError(e, "加密错误")
		}
		return bytes, nil
	}
}
func (b *VlinkCommonBlockchainBaseConfiguration) Decrypt(encData []byte, version uint64) ([]byte, error2.IVlinkError) {
	if key, exist := b.MultiAESSecrets[version]; !exist {
		return nil, error2.NewConfigError(nil, "配置错误,无法找到匹配的密钥")
	} else {
		bytes, e := encrypt.AesDecrypt(encData, []byte(key))
		if nil != e {
			return nil, error2.NewArguError(e, "解密错误")
		}
		return bytes, nil
	}
}
func (c *VlinkCommonBlockchainBaseConfiguration) Config() error2.IVlinkError {
	l1, l2, l3 := len(c.methods), len(c.argsParameters), len(c.baseTypes)
	if l1 != l2 || l2 != l3 || l1 != l3 {
		return error2.NewConfigError(nil, "method和参数转换以及base描述长度必须一致")
	}

	c.ArgsCheckMap = make(map[base.MethodName]*ArgsParameter)
	c.LogicDescriptionMap = make(map[base.MethodName]base.TransBaseDescription)
	for i := 0; i < l1; i++ {
		c.ArgsCheckMap[c.methods[i]] = &c.argsParameters[i]
		c.LogicDescriptionMap[c.methods[i]] = c.baseTypes[i]
	}

	l4, l5 := len(c.keyConstants), len(c.keyGenerators)
	if l4 != l5 {
		return error2.NewConfigError(nil, "keyConstants和keyGenerators必须一致")
	}
	c.BlockChainKeyContainer = make(map[base.ObjectType]base.KeyGenerater)
	for i := 0; i < l4; i++ {
		fmt.Println(fmt.Sprintf("为key=[%v]的注册生成key函数 \n", c.keyConstants[i]))
		c.BlockChainKeyContainer[c.keyConstants[i]] = c.keyGenerators[i]
	}

	c.MultiAESSecrets = make(map[uint64]string)
	c.MultiAESSecrets[0] = "321423u9y8d2fwfl"

	c.methods = nil
	c.baseTypes = nil
	c.argsParameters = nil
	c.keyConstants = nil
	c.keyGenerators = nil

	return nil
}
func (c *VlinkCommonBlockchainBaseConfiguration) GetKey(stub shim.ChaincodeStubInterface, key base.ObjectType, req ...interface{}) (string, error2.IVlinkError) {
	if g, exist := c.BlockChainKeyContainer[key]; !exist {
		return "", error2.NewConfigError(nil, fmt.Sprintf("key=[%v]的生成key函数不存在", key))
	} else {
		return g(stub, key, req...)
	}
}

func NewVlinkCommonBlockchainBaseConfiguration(methods []base.MethodName, args []ArgsParameter, descs []base.TransBaseDescription, decrypter ArgsDecrypter, keyConstants []base.ObjectType, keyGenerators []base.KeyGenerater) *VlinkCommonBlockchainBaseConfiguration {
	c := new(VlinkCommonBlockchainBaseConfiguration)
	c.argsParameters = args
	c.methods = methods
	c.baseTypes = descs
	c.ArgsDecrypter = decrypter
	c.keyConstants = keyConstants
	c.keyGenerators = keyGenerators

	return c
}
