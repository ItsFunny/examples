/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-15 11:48 
# @File : config.go
# @Description : 
# @Attention : 
*/
package service

import (
	"examples/blockchain/config/common/base/fabric"
	"examples/blockchain/config/common/base/fabric/models"
	error2 "examples/blockchain/config/common/error"
)

type IVlinkBlockChainBaseConfiger interface {
	Config() error2.IVlinkError
	CheckAndConvt(method base.MethodName, args []string) (models.BaseFabricAfterValidModel, error2.IVlinkError)
	// ArgsHelper
	// TransactionHelper
}
