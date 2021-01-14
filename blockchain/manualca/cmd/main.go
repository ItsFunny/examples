/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-06 15:07 
# @File : main.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	bcConfig "myLibrary/go-library/blockchain"
)

var (
	configuration *DemoBlockChainConfiguration
)

func init() {
	configuration = NewDemoBlockChainConfiguration()
}

type DemoBlockChainConfiguration struct {
	*bcConfig.BlockChainConfiguration
}

func NewDemoBlockChainConfiguration() *DemoBlockChainConfiguration {
	d := &DemoBlockChainConfiguration{}
	d.BlockChainConfiguration = bcConfig.NewBlockChainConfiguration()

	return d
}

func main() {
	wrapper := bcConfig.ConfigWrapper{}
	err := configuration.Config("/Users/joker/go/src/examples/blockchain/manualca/cmd/application-blockchain-local.yaml", wrapper)
	if nil != err {
		panic(err)
	}

	req := bcConfig.ExecuteReq{
		MethodName:     "asd",
		ChannelID:      "demochannel",
		OrganizationID: "Org0MSP",
		ChainCodeID:    "democc",
		ReqData:        "123",
		Context:        nil,
	}
	resp, bytes, baseError := configuration.Execute(req)
	fmt.Println(baseError)
	fmt.Println(string(bytes))
	fmt.Println(resp)

	engine := gin.Default()

	engine.Run()

}
