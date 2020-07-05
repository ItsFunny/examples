/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-15 12:38 
# @File : blockchain_status.go
# @Description :   链上状态常量
# @Attention : 
*/
package constants

const (
	//  未上链
	NEED_UPLOAD=1
	// 在链
	ON_CHAIN=NEED_UPLOAD<<1

)