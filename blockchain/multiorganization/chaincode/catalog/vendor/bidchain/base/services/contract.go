/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/27 10:39
# @File : contract.go
# @Description :
# @Attention :
*/
package services

import (
	"bidchain/base/config"
	"bidchain/base/models"
	"bidchain/base/services/wrapper"
	"bidchain/fabric/bserror"
	"bidchain/fabric/chaincode/ibidchain_contract"
	"bidchain/fabric/context"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type IContractAdapter interface {
	IBaseService
	ibidchain_contract.IBidchainContract
	// 跨链查询
	InvokeOtherChain(req models.InvokeOtherChainReq) (models.InvokeOtherChainResp, bserror.IBSError)
	// 生成key
	GenCompositeKey(objectType config.ObjectType, args ...interface{}) (string, bserror.IBSError)
	// 获取跨链配置
	GetOverChainBaseInfo(key config.OverChainKey, args []byte) config.OverChainStruct
	// 获取context
	GetContext() context.IBidchainContext
	CheckExist(ot config.ObjectType, keyArgs ...interface{}) (wrapper.KeyWrapper, bool, bserror.IBSError)
	// protoBuf的形式上传数据
	PutInterfaceData(ot config.ObjectType,data interface{},keyArgs ...interface{})bserror.IBSError

	GetByBuildKey(ot config.ObjectType,handler func([]byte)error,keyArgs ...interface{})(string,bserror.IBSError)
	GetByKey(key string,handler func([]byte)error)bserror.IBSError

	PutByKey(key string,data interface{})bserror.IBSError

	GetByPrefix(objectType config.ObjectType,handler func(shim.StateQueryIteratorInterface)error, args []string) bserror.IBSError
}
