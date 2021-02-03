/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/28 17:09
# @File : result.go
# @Description :
# @Attention :
*/
package dto

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ResultDTO struct {
	// code 0 为成功
	Code int64
	Msg  string
	Data []byte
}

func (this ResultDTO) Result() peer.Response {
	bytes, e := json.Marshal(this)
	if nil != e {
		return shim.Error("序列化失败:" + e.Error())
	}
	return shim.Success(bytes)
}

func Fail(code int64, msg string) ResultDTO {
	return ResultDTO{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func EmptySuccess() ResultDTO {
	return ResultDTO{
		Code: 0,
		Msg:  "success",
	}

}

func Success(dataInterface interface{}) ResultDTO {
	var data []byte
	if dataInterface != nil {
		switch dataInterface.(type) {
		case []byte:
			data = dataInterface.([]byte)
		default:
			bytes, e := json.Marshal(dataInterface)
			if nil != e {
				return ResultDTO{
					Code: 1,
					Msg:  "序列化失败:" + e.Error(),
				}
			}
			data = bytes
		}

	}

	return ResultDTO{
		Code: 0,
		Msg:  "success",
		Data: data,
	}

}
