/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-14 16:25 
# @File : result.go
# @Description : 
# @Attention : 
*/
package dto

type ResultDTO struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}
