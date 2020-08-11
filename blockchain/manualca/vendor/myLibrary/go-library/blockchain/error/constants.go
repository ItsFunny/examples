/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-08 14:02 
# @File : constants.go
# @Description : 
# @Attention : 
*/
package error

import "myLibrary/go-library/common/error"

const (
	// fabric调用错误
	FABRIC_ERROR_CODE = 1<<8 | error.FAIL
	// 跨链调用失败
	OVER_CHAINCODE_INVOKE_ERROR = 1<<13 |error.FAIL
	// 账本错误
	FABRIC_LEDGER_ERROR = 1<<16 | error.FAIL


	USER_REGISTRATION_ERROR=1<<18|error.FAIL

	CHANNEL_CONNECTION_ERROR=1<<19|error.FAIL

)