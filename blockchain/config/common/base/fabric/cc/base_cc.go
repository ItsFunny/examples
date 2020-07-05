/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 13:41 
# @File : base.go
# @Description : 
# @Attention : 
*/
package cc

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"examples/blockchain/config/common/base/fabric"
	"examples/blockchain/config/common/base/service/impl"
	error2 "examples/blockchain/config/common/error"
)

type IVlinkChaincode interface {
	shim.Chaincode
	Config() error2.IVlinkError
}

type IConcreteChainCode interface {
	InitDetail(stub shim.ChaincodeStubInterface) error2.IVlinkError
	InvokeDetail(stub shim.ChaincodeStubInterface) base.VlinkPeerResponse
	ConfigDetail() error2.IVlinkError
}

type VlinkBaseChainCodeImpl struct {
	*impl.VlinkBaseServiceImpl
	ConcreteChainCode IConcreteChainCode
}

func NewVlinkBaseChainCodeImpl(ConcreteChainCode IConcreteChainCode) *VlinkBaseChainCodeImpl {
	v := new(VlinkBaseChainCodeImpl)
	v.VlinkBaseServiceImpl = impl.NewVlinkBaseServiceImpl()
	v.ConcreteChainCode = ConcreteChainCode

	return v
}

// 启动时候的调用,发生在init之前,当跨链
// func (v *VlinkBaseChainCodeImpl) BootStrapConfig() peer.Response {
// }

func (v *VlinkBaseChainCodeImpl) Init(stub shim.ChaincodeStubInterface) peer.Response {
	v.VlinkBaseServiceImpl = impl.NewVlinkBaseServiceImpl()
	e := v.ConcreteChainCode.InitDetail(stub)
	if nil != e {
		return shim.Error(e.Error())
	}
	e = v.ConcreteChainCode.ConfigDetail()
	if nil != e {
		return shim.Error(e.Error())
	}

	return shim.Success(nil)
}

func (v *VlinkBaseChainCodeImpl) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	resp := v.ConcreteChainCode.InvokeDetail(stub)
	return resp.Response
}

func (v *VlinkBaseChainCodeImpl) Config() error2.IVlinkError {
	return v.ConcreteChainCode.ConfigDetail()
}

// type IChaincodeFacadedService interface {
// 	log.Logger
// 	ValidateArguAndReturn(method base.MethodName, args []string) (models.BaseFabricAfterValidModel, error2.IVlinkError)
// }

// 参数加解密

// type IArgumentDecrypt interface {
// 	Decrypt(argu interface{}, version string) (interface{}, error2.IVlinkError)
// 	SetParent(c IChaincodeFacadedService)
// }
