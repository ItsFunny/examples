/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-28 12:53 
# @File : baseAPIModel.go
# @Description : 
*/
package models

type BaseAPIResponseModel struct {
	Msg  string
	Code string
}

func (this *BaseAPIResponseModel) SetResponseCode(code string) {
	this.Code = code
}

func (this *BaseAPIResponseModel) GetResponseCode() string {
	return this.Code
}

func (this *BaseAPIResponseModel) SetResponseMsg(msg string) {
	this.Msg = msg
}

func (this *BaseAPIResponseModel) GetResponseMsg() string {
	return this.Msg
}


