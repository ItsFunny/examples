/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-25 16:15 
# @File : config.go
# @Description : 
# @Attention : 
*/
package config

import (
	"fmt"
	"myLibrary/go-library/go/blockchain"
	"myLibrary/go-library/go/blockchain/base"
	"myLibrary/go-library/go/blockchain/model"
)

var (
	configuration *DemoConfiguration
)

func init() {
	configuration = new(DemoConfiguration)
	configuration.BlockChainConfiguration = config.NewBlockChainConfiguration()
}

type DemoConfiguration struct {
	*config.BlockChainConfiguration
}

func Config(path string) error {
	return configuration.Config(path)
}

func Enrolle(userName string, pwd string) error {
	return configuration.Enroll(model.UserEnrollReq{
		Oid:           "Org1MSP",
		UserUniqueKey: userName,
		UserPassword:  pwd,
		Profile:       "",
		Type:          "admin",
	})
}

func Registraiton(userName string, pwd string) error {
	resp, baseError := configuration.Register(model.UserRegisterReq{
		Oid:    "Org1MSP",
		Name:   userName,
		Secret: pwd,
		Type:   "client",
	})
	fmt.Println(resp)
	return baseError
}
func Invoke() {
	resp, bytes, baseError := configuration.Execute(config.ExecuteReq{
		MethodName:     "asd",
		ChannelID:      "demochannel",
		OrganizationID: "Org1MSP",
		ChainCodeID:    "democc",
		ReqData:        "asddd",
	})
	fmt.Println(baseError)
	fmt.Println(string(bytes))
	fmt.Println(resp)
}

func NewAndInvoke() {
	client, e := configuration.NewChannelClient("demochannel", "Org1MSP", "joker")
	if nil != e {
		panic(e)
	}
	req := base.ChainBaseReq{
		MethodName:     "asdddd",
		ChannelID:      "demochannel",
		OrganizationID: "Org1MSP",
		ChainCodeID:    "democc",
	}
	withClient, baseError := configuration.ExecuteWithClient(client, req, "asdd")
	fmt.Println(baseError)
	fmt.Println(withClient)
}
