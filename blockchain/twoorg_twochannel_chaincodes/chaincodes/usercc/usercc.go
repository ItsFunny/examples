/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-03 13:31 
# @File : usercc.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type UserCC struct {
}

func (c *UserCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("usercc init")
	return shim.Success(nil)
}

func (c *UserCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	panic("implement me")
}

func main() {
	u := new(UserCC)
	if e := shim.Start(u); nil != e {
		panic(e)
	}
}
