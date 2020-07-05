/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-03-26 13:23 
# @File : demo_chaincode.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type DemoChainCode struct {
}

func (this *DemoChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("init")
	return shim.Success(nil)
}

func (this *DemoChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	for _, str := range args {
		fmt.Println(str)
	}

	return shim.Success([]byte("success"))
}

func main() {
	u := new(DemoChainCode)
	if err := shim.Start(u); nil != err {
		panic(err)
	}
}
