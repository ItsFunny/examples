/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:58 
# @File : base_blockchain.go
# @Description : 
# @Attention : 
*/
package impl

import (
	"encoding/json"
	"myLibrary/go-library/go/base/models"
	"myLibrary/go-library/go/converters"
	error2 "myLibrary/go-library/go/error"
	"examples/blockchain/config/common/base/fabric"
)

// Deprecated: use BaseFabricResp instead
type VlinkBaseBlockChainResp struct {
	*models.BaseResponseModel
}

func NewVlinkBaseBlockChainResp() *VlinkBaseBlockChainResp {
	b := new(VlinkBaseBlockChainResp)
	b.BaseResponseModel = models.NewBaseResponseModel()

	return b
}

func SuccessWithEmptyData() []byte {
	bytes, _ := json.Marshal(*Success())
	return bytes
}
func SuccessWithEmptyDataWithDetail(code int, msg string) []byte {
	r := new(base.BaseFabricResp)
	r.CodeBytes = converter.BigEndianInt642Bytes(int64(code))
	r.MsgBytes = []byte(msg)
	bytes, _ := json.Marshal(r)
	return bytes
}

func FailWithBytes(code error2.ErrorCode, msg string) []byte {
	m := base.BaseFabricResp{
		DataBytes: nil,
		CodeBytes: converter.BigEndianInt642Bytes(int64(code)),
		MsgBytes:  []byte(msg),
	}
	bytes, _ := json.Marshal(m)

	return bytes
}
func Success() *base.BaseFabricResp {
	resp := base.NewSuccessBaseFabricResp()
	return resp
}

func Fail(code int, msg string) *VlinkBaseBlockChainResp {
	resp := NewVlinkBaseBlockChainResp()
	resp.Code = code
	resp.Msg = msg
	return resp
}
