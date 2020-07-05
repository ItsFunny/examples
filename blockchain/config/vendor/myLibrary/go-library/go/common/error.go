/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-20 11:12 
# @File : error.go
# @Description : 
*/
package common

type BaseError struct {
	Msg string
}

func (this BaseError) Error() string {
	return this.Msg
}

type DBError struct {
	BaseError
}

type FabricError struct {
	BaseError
}
type SystemError struct {
	BaseError
}
type BussError struct {
	BaseError
}

func NewDbErr(err error) DBError {
	return DBError{
		BaseError: BaseError{
			Msg: err.Error(),
		},
	}
}

func NewSysErr(err error) SystemError {
	return SystemError{
		BaseError: BaseError{
			Msg: err.Error(),
		},
	}
}
func NewBussErr(err error) BussError {
	return BussError{
		BaseError: BaseError{
			Msg: err.Error(),
		},
	}
}
func NewFabricErr(err error) FabricError {
	return FabricError{
		BaseError: BaseError{
			Msg: err.Error(),
		},
	}
}