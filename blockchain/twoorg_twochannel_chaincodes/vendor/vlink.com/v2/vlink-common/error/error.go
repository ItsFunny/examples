/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:03 
# @File : error.go
# @Description : error封装类,通过
# @Attention : 
*/
package error

import (
	error2 "myLibrary/go-libary/go/error"
	"vlink.com/v2/vlink-common/constants"
)

type IVlinkError interface {
	error2.IBaseError
}

type VlinkError struct {
	*error2.BaseError
}

// 有效
func (v *VlinkError) Valid() bool {
	return v.Code != 0
}

func ErrorsWithDetail(err error, code error2.ErrorCode, msg string) IVlinkError {
	if nil != err {
		return nil
	}
	return VlinkError{
		BaseError: error2.NewBaseError(err, code, msg),
	}
}

func ErrorsWithMessage(err IVlinkError, msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(err, constants.ARGUMENT_ERROR_CODE, msg),
	}
}

func NewArguError(e error, msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(e, constants.ARGUMENT_ERROR_CODE, msg),
	}
}
func NewJSONSerializeError(e error, msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(e, constants.JSON_SERIALIZE_ERROR_CODE, msg),
	}
}
func NewJSONUnSerializeError(e error, msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(e, constants.JSON_UN_SERIALIZE_ERROR_CODE, msg),
	}
}

func NewConfigError(e error, msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(e, constants.CONFIG_ERROR_CODE, msg),
	}
}

func NewFabricError(e error, msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(e, constants.FABRIC_ERROR_CODE, msg),
	}
}

func NewSystemError(e error, msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(e, constants.SYSTEM_ERROR_CODE, msg),
	}
}

func NewLogicOutError(msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(nil, constants.OUT_PUT_ERROR_CODE, msg),
	}
}
