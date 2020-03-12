/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 13:43 
# @File : base.go
# @Description : 
# @Attention : 
*/
package error

type ErrorCode int

type IBaseError interface {
	error
	GetCode() ErrorCode
	SetCode(code ErrorCode)
	GetMsg() string
	SetMsg(m string)
}
