/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-18 14:02 
# @File : wallet.go
# @Description : 
# @Attention : 
*/
package wallet

const (
	WALLET_CHANGE_IN  = 0
	WALLET_CHANGE_OUT = 1
)

const (
	// 主钱包
	WALLET_TYPE_MAIN_WALLET = 1

	// 记录交易的子钱包
	WALLET_TYPE_CHILD_RECORD_TX = WALLET_TYPE_MAIN_WALLET << 1
)
