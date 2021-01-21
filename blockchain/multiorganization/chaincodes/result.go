/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-20 15:37 
# @File : result.go
# @Description : 
# @Attention : 
*/
package main

import (
	"encoding/json"
)


const (
	OK = iota + 1
	FAIL
	DATA_NOT_EXIST


	HISTORY_ITERATOR_FAIL=100
	HISTORY_EMPTY = 101
	HISTORY_DOESNT_HAS_NEXT=102
	HISTORY_ERROR=103
	HISTORY_CONTINUE_FIND=1000
	HISTORY_FINALLY_FOUND_AFTER_SEVERALTIMES=10000
)


type ResultDTO struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func DefaultWith(code int, msg string, data []byte) []byte{
	dto := create(code, msg, data, )
	bytes, _ := json.Marshal(dto)
	return bytes
}
func DefaultSuccessWith(data []byte)[]byte{
	return SuccessWith(data, func(dto *ResultDTO) []byte {
		bytes, _ := json.Marshal(dto)
		return bytes
	})
}


func SuccessWith(data []byte, f func(dto *ResultDTO) []byte) []byte {
	success := Success(data)
	return f(success)
}

func Success(data []byte) *ResultDTO {
	return &ResultDTO{
		Code: OK,
		Msg:  "",
		Data: string(data),
	}
}

func create(code int, msg string, data []byte) *ResultDTO {
	return &ResultDTO{
		Code: code,
		Msg:  msg,
		Data: string(data),
	}
}

func Fail(code int, msg string, data []byte) *ResultDTO {
	return create(code, msg, data)
}
