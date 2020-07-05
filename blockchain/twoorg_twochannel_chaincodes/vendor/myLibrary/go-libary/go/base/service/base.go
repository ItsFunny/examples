/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 14:15 
# @File : base_service.go
# @Description : 
# @Attention : 
*/
package service

import "myLibrary/go-libary/go/log"

type IBaseService interface {
	// 支持日志
	log.Logger
	BeforeStart(methodName string)
	AfterEnd()
}
