/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-04 09:21 
# @File : tx_desc.go
# @Description : 
# @Attention : 
*/
package txbase

import "examples/blockchain/config/common/base/fabric"

const (
	// 更新用户积分
	COINCC_TX_BASE_USER_COIN_UPDATE = (base.TX_BASE_NEED_RECORD << 1) | base.TX_BASE_NEED_RECORD
	// 获取用户积分,不需要记录
	COINCC_TX_BASE_USER_COIN_GET=base.TX_BASE_UNNEED_RECORD<<1
)
