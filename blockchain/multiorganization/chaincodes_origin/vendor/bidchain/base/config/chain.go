/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/28 16:34
# @File : chain.go
# @Description :
# @Attention :
*/
package config

type OverChainStruct struct {
	ChannelName   string
	ChaincodeName string
	MethodName    string
	Args          []byte
}
