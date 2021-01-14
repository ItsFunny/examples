/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-23 14:56 
# @File : UserFacadedCC.go
# @Description : 
*/
package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type UserFacadedCC struct {
}

func (u *UserFacadedCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("user-init")
	return shim.Success(nil)
}

func (u *UserFacadedCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	methodName, strings := stub.GetFunctionAndParameters()
	str := ""
	for _, s := range strings {
		str += s
	}
	defer func() {
		e := stub.SetEvent("123", []byte("success"))
		if nil != e {
			fmt.Println("注册event失败")
		}
	}()
	fmt.Println("参数为:", str)
	switch methodName {
	case "fromCC":
		return shim.Success([]byte("这是user的cc"))
	case "userPayInvoke":
		res := stub.InvokeChaincode("paycc", [][]byte{[]byte("fromCC"), []byte("hello")}, "")
		return shim.Success([]byte("跨链调用pay的Invoke:" + string(res.Payload)))
	default:
		fmt.Println("usercc-default")
		return shim.Success([]byte("user-invoke"))
	}
}

func main() {
	e := shim.Start(new(UserFacadedCC))
	if nil != e {
		fmt.Println("start user cc failed for reason :", e.Error())
		panic(e)
	}
}
