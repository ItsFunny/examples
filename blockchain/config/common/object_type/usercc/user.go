/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-16 10:09 
# @File : user.go
# @Description : 用户相关的compositeKey,以OT(objectType)开头
# @Attention : 
*/
package usercc

import "vlink.com/v2/vlink-common/base/fabric"

const (
	// 用户ID和DNA的映射
	OT_USER_ID_DNA = base.ObjectType("OT_USER_ID_DNA")
	// 用户钱包
	OT_USER_WALLET = base.ObjectType("OT_USER_WALLET")
	// 用户信息
	OT_USER_INFO = base.ObjectType("OT_USER_INFO")
)
