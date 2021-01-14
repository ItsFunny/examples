/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:55 
# @File : base.go
# @Description : 
# @Attention : 
*/
package models

type IBaseResponse interface {
	GetCode() int
	SetCode(code int ) int
	GetMsg() string
	SetMsg(msg string)
}
