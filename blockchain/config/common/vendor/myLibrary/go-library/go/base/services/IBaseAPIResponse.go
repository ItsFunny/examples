/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-28 12:53 
# @File : IBaseAPIResponse.go
# @Description : 
*/
package services


type IBaseRepsonseService interface {
	SetResponseCode(code string)
	GetResponseCode()string
	SetResponseMsg(msg string)
	GetResponseMsg()string
}
