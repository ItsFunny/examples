/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:56 
# @File : base_response.go
# @Description : 
# @Attention : 
*/
package models

type BaseResponseModel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewBaseResponseModel() *BaseResponseModel {
	b := new(BaseResponseModel)
	return b
}

func (b *BaseResponseModel) GetCode() int {
	return b.Code
}

func (b *BaseResponseModel) SetCode(code int) int {
	b.Code = code
	return code
}

func (b *BaseResponseModel) GetMsg() string {
	return b.Msg
}

func (b *BaseResponseModel) SetMsg(msg string) {
	b.Msg = msg
}
