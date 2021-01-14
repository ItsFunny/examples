/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-23 14:58 
# @File : ItemFacadedCC.go
# @Description : 
*/
package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ItemFacadedCC struct {
}

func (u *ItemFacadedCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("item-init")
	return shim.Success(nil)
}

func (u *ItemFacadedCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	methodName, strings := stub.GetFunctionAndParameters()
	str := ""
	for _, s := range strings {
		str += s
	}
	fmt.Println("参数为:", str)
	switch methodName {
	case "fromCC":
		return shim.Success([]byte("这是item的cc"))
	case "itemUserInvoke":
		res := stub.InvokeChaincode("fromCC", [][]byte{[]byte("fromCC"), []byte("hello")}, "")
		return shim.Success([]byte("跨链调用user的Invoke:" + string(res.Payload)))
	default:
		fmt.Println("itemcc-default")
		return shim.Success([]byte("user-invoke"))
	}
}

func main() {
	e := shim.Start(new(ItemFacadedCC))
	if nil != e {
		fmt.Println("start item cc failed for reason :", e.Error())
		panic(e)
	}
}
