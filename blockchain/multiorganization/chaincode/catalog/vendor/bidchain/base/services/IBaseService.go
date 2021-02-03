/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-19 12:17 
# @File : IBaseService.go
# @Description : 
*/
package services

import (
	"bidchain/fabric/log"
)

type (
	IBaseServiceInit interface {
		GetReqId() string
		GetLogger() log.Logger
		SetReqId(string)
		SetLogger(log.Logger)
	}

	IBaseService interface {
		GetInitInfo() IBaseServiceInit
		SetInitInfo(init IBaseServiceInit)
		BeforeStart(method string)
		AfterEnd()
	}

	IParamVerifier interface {
		Validation() error
	}
)

