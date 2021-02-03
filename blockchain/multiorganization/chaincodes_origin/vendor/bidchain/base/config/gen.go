/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/27 15:37
# @File : gen.go
# @Description :
# @Attention :
*/
package config

import (
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type KeyGenerator func(stub shim.ChaincodeStubInterface, objectType ObjectType, params ...interface{}) (string, error)

var (
	COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string, error) {
		switch param[0].(type) {
		case string:
			return COMMON_STRING_KEY_GENERATOR(stub, objectType, param[0])
		case []string:
			return COMMON_STRING_ARRAY_KEY_GENERATOR(stub, objectType, param[0])
		}
		return "", errors.New("找不到匹配的处理")
	}
	COMMON_STRING_ARRAY_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string, error) {
		strings := param[0].([]string)
		s, e := stub.CreateCompositeKey(string(objectType), strings)
		if nil != e {
			return "", errors.New("创建组合键失败:" + e.Error())
		}
		return s, nil
	}

	COMMON_STRING_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string, error) {
		strings := param[0].(string)
		s, e := stub.CreateCompositeKey(string(objectType), []string{strings})
		if nil != e {
			return "", errors.New("创建组合键失败:" + e.Error())
		}
		return s, nil
	}
)
