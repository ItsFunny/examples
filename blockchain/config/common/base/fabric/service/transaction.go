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
	"examples/blockchain/config/common/base/fabric"
	error2 "examples/blockchain/config/common/error"
)

type TransactionHelper interface {
	GetBaseType() error2.IVlinkError
}

type IBaseTxService interface {
	GetBCBase() base.BCBase
}
