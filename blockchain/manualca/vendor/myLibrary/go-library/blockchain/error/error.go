/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-08 13:59 
# @File : error.go
# @Description : 
# @Attention : 
*/
package error

import error2 "myLibrary/go-library/common/error"

type BlockChainError struct {
	*error2.BaseError
}
