/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-18 16:21 
# @File : vlink_configuration.go
# @Description : 
# @Attention : 
*/
package config

import (
	"examples/blockchain/config/blockchain"
	error2 "examples/blockchain/config/common/error"
	"fmt"
)

type VlinkConfiguration struct {
	VlinkBlockChainConfiguration *blockchain.VlinkBlockChainConfiguration
	Properties *VlinkProperties
}

type VlinkProperties struct {
	VlinkBlockChainProperties blockchain.VlinkBlockChainProperties `yaml:"blockchain" json:"vlinkBlockChainProperties"`
}

func (p *VlinkProperties) config() error2.IVlinkError {
	fmt.Println("begin 初始化 区块链")
	// 初始化blockchain
	e := configuration.VlinkBlockChainConfiguration.Config(p.VlinkBlockChainProperties)
	if nil != e {
		fmt.Println("配置区块链失败:", e.Error())
		return error2.NewConfigError(e, "配置区块链失败")
	}
	fmt.Println("end 初始化 区块链")
	return nil
}
