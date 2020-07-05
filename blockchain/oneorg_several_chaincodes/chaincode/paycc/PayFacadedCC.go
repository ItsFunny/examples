/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-23 14:55 
# @File : PayFacadedCC.go
# @Description : 
*/
package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type PayFacadedCC struct {
}

func (p *PayFacadedCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("pay-cc init ")
	return shim.Success(nil)
}

func (p *PayFacadedCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	methodName, strings := stub.GetFunctionAndParameters()
	str := ""
	for _, s := range strings {
		str += s
	}
	fmt.Println("参数为:", str)
	switch methodName {
	case "fromCC":
		return shim.Success([]byte("这是pay的cc"))
	case "payItemInvoke":
		res := stub.InvokeChaincode("itemcc", [][]byte{[]byte("fromCC"), []byte("hello")}, "")
		return shim.Success([]byte("跨链调用item的Invoke:" + string(res.Payload)))
	default:
		fmt.Println("paycc-default")
		return shim.Success([]byte("pay-invoke"))
	}
}

func main() {
	e := shim.Start(new(PayFacadedCC))
	if nil != e {
		fmt.Println("start pay cc failed for reason :", e.Error())
		panic(e)
	}
}
