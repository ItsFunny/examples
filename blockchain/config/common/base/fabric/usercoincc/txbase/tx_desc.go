/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-07 09:48 
# @File : tx_desc.go
# @Description : 
# @Attention : 
*/
package txbase

import "vlink.com/v2/vlink-common/base/fabric"

const (
	USERCOIN_TX_BASE_UPLOAD_USER_WITH_INIT_COIN = (base.TX_BASE_NEED_RECORD << 1) | base.TX_BASE_NEED_RECORD
	USERCOIN_TX_BASE_UPDATE_USER_INFO           = USERCOIN_TX_BASE_UPLOAD_USER_WITH_INIT_COIN << 1
)
