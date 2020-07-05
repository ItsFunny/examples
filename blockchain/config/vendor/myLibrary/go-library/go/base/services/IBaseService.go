/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-19 12:17 
# @File : IBaseService.go
# @Description : 
*/
package services

import (
	"myLibrary/go-library/go/common/log"
)

type (
	IBaseServiceInit interface {
		GetReqId() string
		GetLogger() *log.Log
		SetReqId(string)
		SetLogger(*log.Log)
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

