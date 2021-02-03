/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/27 10:38
# @File : invoke.go
# @Description :
# @Attention :
*/
package models

type InvokeOtherChainReq struct {
	ChannelName   string
	ChaincodeName string
	MethodName    string
	// args[0] 为方法 名称 , args[1] 为参数
	Args []byte
}

func (i InvokeOtherChainReq) String() string {
	return "channelName:" + i.ChannelName + ",chaincodeName:" + i.ChaincodeName
}

type InvokeOtherChainResp struct {
	Payload []byte
}
