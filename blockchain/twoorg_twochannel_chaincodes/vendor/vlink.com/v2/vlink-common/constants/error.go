/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:04 
# @File : error.go
# @Description : 错误异常的常量
# @Attention : 
*/
package constants

const (
	// 成功
	SUCCESS=1
	// 需要显示描述的错误,如用户重复上链等信息
	OUT_PUT_ERROR_CODE = 1 << 1



	FAIL = 5 << 1
	// 系统错误
	SYSTEM_ERROR_CODE = FAIL<<1 | FAIL
	// 数据库错误
	DB_ERROR_CODE = SYSTEM_ERROR_CODE<<1 | FAIL
	// fabric调用错误
	FABRIC_ERROR_CODE = DB_ERROR_CODE<<1 | FAIL
	// 参数错误
	ARGUMENT_ERROR_CODE = FABRIC_ERROR_CODE<<1 | FAIL
	// 配置错误
	CONFIG_ERROR_CODE = ARGUMENT_ERROR_CODE<<1 | FAIL
	// json序列化错误
	JSON_SERIALIZE_ERROR_CODE    = CONFIG_ERROR_CODE<<1 | FAIL
	JSON_UN_SERIALIZE_ERROR_CODE = JSON_SERIALIZE_ERROR_CODE<<1 | FAIL
	// 系统错误
	// 需要外抛的错误
	RETURN_ERROR_CODE = SYSTEM_ERROR_CODE<<62 | FAIL
)
