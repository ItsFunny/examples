/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/28 17:45
# @File : error_code.go
# @Description :
# @Attention :
*/
package bserror

type ErrorCode int64

var (
	BAD_REQUEST_ERROR_CODE int64 = 400
	BAD_REQUEST_NO_AUTH    int64 = 401
)
