/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/27 16:29
# @File : error_adapter'.go
# @Description :
# @Attention :
*/
package bserror

type IBSError interface {
	error
	SetMsg(msg string)
	GetMsg()string
	SetCode(errorCode int64)
	GetCode()int64
}

