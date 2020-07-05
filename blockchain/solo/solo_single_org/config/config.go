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
	"myLibrary/go-library/go/blockchain"
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

func Enrolle(userName string,pwd string)error{
	return configuration.Enroll(model.UserEnrollReq{
		Oid:           "Org1MSP",
		UserUniqueKey: userName,
		UserPassword:  pwd,
		Profile:       "",
		Type:          "",
	})
}