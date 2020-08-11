/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:04 
# @File : error.go
# @Description : 结果 错误异常的常量
# @Attention : 
*/
package constants

const (
	// 成功
	SUCCESS = 1
	// 需要显示描述的错误,如用户重复上链等信息
	OUT_PUT_ERROR_CODE = 1 << 1

	FAIL = 5 << 1
	// 系统错误
	SYSTEM_ERROR_CODE = 1<<6 | FAIL
	// 数据库错误
	DB_ERROR_CODE = 1<<7 | FAIL


	// 参数错误
	ARGUMENT_ERROR_CODE = 1<<9 | FAIL
	// 配置错误
	CONFIG_ERROR_CODE = 1<<10 | FAIL
	// json序列化错误
	JSON_SERIALIZE_ERROR_CODE    = 1<<11 | FAIL
	JSON_UN_SERIALIZE_ERROR_CODE = 1<<12 | FAIL


	// http失败
	HTTP_NETWORK_ERROR = 1<<14 | FAIL

	SPIDE_ERROR = 1<<15 | FAIL


	CRYPT_ERROR         = 1<<17 | FAIL
	// 系统错误

	// 需要外抛的错误
	RETURN_ERROR_CODE = 1<<32 | FAIL
	OUTPUT_ERROR_CODE = 1<<33 | RETURN_ERROR_CODE
	RECORD_NOT_EXIST_ERROR = 1<<34 | OUTPUT_ERROR_CODE
)
