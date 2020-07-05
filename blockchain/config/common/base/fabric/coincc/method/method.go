/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-05 10:08 
# @File : method.go
# @Description : 
# @Attention : 
*/
package method

import "examples/blockchain/config/common/base/fabric"

const (
	// 更新用户积分
	COINCC_METHOD_USER_COIN_UPDATE = base.MethodName("UpdateUserCoin")

	// 获取用户积分
	COINCC_METHOD_USER_COIN_GET=base.MethodName("GetUserCoin")
)
