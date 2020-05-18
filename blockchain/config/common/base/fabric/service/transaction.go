/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-15 11:50 
# @File : transaction.go
# @Description : 
# @Attention : 
*/
package service

import (
	"vlink.com/v2/vlink-common/base/fabric"
	error2 "vlink.com/v2/vlink-common/error"
)

type TransactionHelper interface {
	GetBaseType() error2.IVlinkError
}

type IBaseTxService interface {
	GetBCBase() base.BCBase
}
