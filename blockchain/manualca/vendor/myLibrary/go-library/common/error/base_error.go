/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 13:44 
# @File : base_error.go
# @Description : 
# @Attention : 
*/
package error


type BaseError struct {
	Cause error
	Code  ErrorCode
	Msg   string
}

func NewBaseError(ee error, c ErrorCode, msg string) *BaseError {
	e := new(BaseError)
	e.Code = c
	e.Msg = msg
	e.Cause = ee

	return e
}
func NewJSONSerializeError(e error, msg string) *BaseError {
	return NewBaseError(e, JSON_SERIALIZE_ERROR_CODE, msg)
}

func NewArguError(e error, msg string) IBaseError {
	return NewBaseError(e, ARGUMENT_ERROR_CODE, msg)
}

type BlockChainLedgerError struct {
	*BaseError
	ChannelID string
}

// 账本错误
func NewLedgerError(e error, id string, msg string) IBaseError {
	er := BlockChainLedgerError{
		BaseError: nil,
		ChannelID: "",
	}
	er.ChannelID = id
	er.BaseError = NewBaseError(e, FABRIC_LEDGER_ERROR, msg)
	return er
}

func NewRecordNotExistError(msg string) IBaseError {
	return NewBaseError(nil, RECORD_NOT_EXIST_ERROR, msg)
}

func NewConfigError(e error, msg string) IBaseError {
	return NewBaseError(e, CONFIG_ERROR_CODE, msg)
}

func NewSystemError(e error, msg string) IBaseError {
	return NewBaseError(e, SYSTEM_ERROR_CODE, msg)
}
func ErrorsWithMessage(err error, msg string) IBaseError {
	return NewBaseError(err, FAIL, msg)
}

func (e *BaseError) Error() string {
	if nil != e.Cause {
		return e.Msg + ":" + e.Cause.Error()
	}

	return e.Msg
}

func (e *BaseError) GetCode() ErrorCode {
	return e.Code
}

func (e *BaseError) SetCode(code ErrorCode) {
	e.Code = code
}

func (e *BaseError) GetMsg() string {
	return e.Msg
}

func (e *BaseError) SetMsg(m string) {
	e.Msg = m
}
