/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-03 17:01 
# @File : User.go
# @Description : 
# @Attention : 
*/
package coincc

import "errors"

type BCUserCoinUpdateReq struct {
	// 用户积分地址
	UserCoinWalletAddress string `json:"userCoinWalletAddress"`
	// 用户要添加的积分数据 ,增加或者减少 ,可以是任意值
	CoinPoint float64 `json:"coinPoint"`
}

func (r BCUserCoinUpdateReq) Validate() error {
	if r.UserCoinWalletAddress == "" {
		return errors.New("参数积分地址不可为空")
	}
	return nil
}

type BCUserCoinUpdateResp struct {
	CoinAmount float64 `json:"coinAmount"`
	// FIXME 需求确定是否要记录 积分的增长过程
}

// 获取用户积分信息
type BCUserCoinGetReq struct {
	// 用户积分地址
	UserCoinWalletAddress string `json:"userCoinWalletAddress"`
}
type BCUserCoinGetResp struct {
	// 积分总额
	Amount float64 `json:"amount"`
}