/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:41 
# @File : log.go
# @Description : 日志的底层提供多种选择,现在仅只提供log4go实现的日志
# @Attention : 
*/
package log

import (
	"myLibrary/go-library/go/log"
	"myLibrary/go-library/go/utils"
)

type VlinkLog struct {
	*log.CommonBaseLogger
}

func NewVlinkLog() *VlinkLog {
	l := new(VlinkLog)
	reqID := utils.GenerateUUID()
	l.CommonBaseLogger = log.NewCommonBaseLoggerWithLog4go(reqID)
	return l
}
