/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-03 13:29 
# @File : coincc.go
# @Description : 
# @Attention : 
*/
package coincc

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CoinCC struct {
}

func (c *CoinCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("init")
	return shim.Success(nil)
}

func (c *CoinCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	panic("implement me")
}

func main() {
	c := new(CoinCC)
	if e := shim.Start(c); nil != e {
		panic(e)
	}
}
