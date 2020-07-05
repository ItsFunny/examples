/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-19 14:15
# @File : config.go
# @Description :
# @Attention :
*/
package config

import (
	"examples/blockchain/config/blockchain"
)

var (
	configuration *VlinkConfiguration
)

func init() {
	configuration = new(VlinkConfiguration)
	if nil == configuration.VlinkBlockChainConfiguration {
		configuration.VlinkBlockChainConfiguration = new(blockchain.VlinkBlockChainConfiguration)
	}
}

func GetBlockChainSetup() *blockchain.VlinkBlockChainConfiguration{
	return configuration.VlinkBlockChainConfiguration
}
