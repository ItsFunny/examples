/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-10 13:54 
# @File : coin.go
# @Description : 
# @Attention : 
*/
package parent



type CoinGetUserCoinAmountReqParent struct {
	UserCoinWalletAddress string `json:"userCoinWalletAddress"`
}

type CoinGetUserCoinAmountRespParent struct {
	// 积分总额
	Amount float64 `json:"amount"`
}