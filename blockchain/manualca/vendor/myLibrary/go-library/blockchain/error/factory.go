/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-08 14:02 
# @File : factory.go
# @Description : 
# @Attention : 
*/
package error

import error2 "myLibrary/go-library/common/error"

func NewFabricError(e error, msg string) error2.IBaseError {
	return error2.NewBaseError(e, FABRIC_ERROR_CODE, msg)
}
func NewChannelError(e error,msg string)error2.IBaseError{
	return error2.NewBaseError(e,CHANNEL_CONNECTION_ERROR,msg)
}
