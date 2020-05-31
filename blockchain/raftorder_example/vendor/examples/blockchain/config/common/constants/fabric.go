/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-19 10:03 
# @File : fabric.go
# @Description : fabric常量,用于指定fabric client从而调用区块链函数 ,全部的channel ,以及chaincode都会存储在这里
# @Attention : value必须与配置文件中的一致
*/
package constants

const (
	// 用户链
	CHANNEL_ID_USER     = "userchannel"
	CHAINCODE_ID_USERCC = "usercc"
	CHAINCODE_ID_USERCOIN="usercoincc"
	// 积分链
	CHANNEL_ID_COIN = "coinchannel"
	// 纯积分链码
	CHAINCODE_ID_COINCC = "coincc"

	// 版权链
	CHANNEL_ID_COPYRIGHT="copyrightchannel"
	// 纯版权积分链码
	CHAINCODE_ID_COPYRIGHTCC="copyrightcc"



)
