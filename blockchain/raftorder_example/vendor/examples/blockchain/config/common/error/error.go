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
	error2 "myLibrary/go-library/go/error"
	"examples/blockchain/config/common/constants"
)

type IVlinkError interface {
	error2.IBaseError
}

type VlinkError struct {
	*error2.BaseError
}

// 链码之间调用失败
type ChainCodeInvokeError struct {
	*VlinkError
	FromChannel string
	ToChannel   string
	MethodName  string
	ChainCodeId string
}

// http 请求错误
type HttpNetWorkError struct {
	*VlinkError
	HttpAddr string
}

// 磁盘错误
type DiskError struct {
	*VlinkError
	Path string
}

type BlockChainLedgerError struct {
	*VlinkError
	ChannelID string
}

type CryptError struct {
	*VlinkError
	// 0 代表加密
	// 1. 代表解密
	Type int
}

// 区块链admin客户端错误
type BlockChainAdminClientError struct {
	*VlinkError
	ChannelID      string
	OrganizationID string
}

// 爬虫爬取错误
type SpideError struct {
	*VlinkError
	Url string
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

func NewCryptError(e error, Type int, msg string) IVlinkError {
	er := CryptError{}
	er.VlinkError = &VlinkError{
		BaseError: error2.NewBaseError(e, constants.CRYPT_ERROR, msg),
	}
	er.Type = Type

	return er
}

// 账本错误
func NewLedgerError(e error, id string, msg string) IVlinkError {
	er := BlockChainLedgerError{
		VlinkError: nil,
		ChannelID:  "",
	}
	er.ChannelID = id
	er.VlinkError = &VlinkError{
		BaseError: error2.NewBaseError(e, constants.FABRIC_LEDGER_ERROR, msg),
	}
	return er
}

func NewAdminClientError(e error, id string, oid string, msg string) IVlinkError {
	er := BlockChainAdminClientError{
		VlinkError:     nil,
		ChannelID:      id,
		OrganizationID: oid,
	}
	er.VlinkError = &VlinkError{
		BaseError: error2.NewBaseError(e, constants.FABRIC_LEDGER_ERROR, msg),
	}
	return er
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

func NewVlinkChainCodeInvokeError(er error, msg string, from, to string, chaincodeId string, methodName string) IVlinkError {
	e := ChainCodeInvokeError{
		FromChannel: from,
		ToChannel:   to,
		MethodName:  methodName,
		ChainCodeId: chaincodeId,
	}
	e.VlinkError = &VlinkError{
		BaseError: error2.NewBaseError(er, constants.OVER_CHAINCODE_INVOKE_ERROR, msg),
	}
	return e
}

func NewHttpNetWorkError(addr string) IVlinkError {
	e := HttpNetWorkError{
		HttpAddr: addr,
	}
	e.VlinkError = &VlinkError{
		BaseError: error2.NewBaseError(nil, constants.HTTP_NETWORK_ERROR, "http调用失败"),
	}
	return e
}

func NewDiskIOError(path string, msg string) IVlinkError {
	e := DiskError{
		Path: path,
	}
	e.VlinkError = &VlinkError{
		BaseError: error2.NewBaseError(nil, constants.HTTP_NETWORK_ERROR, "磁盘调用错误:"+msg),
	}
	return e
}

func NewDBError(msg string) IVlinkError {
	return VlinkError{
		BaseError: error2.NewBaseError(nil, constants.DB_ERROR_CODE, msg),
	}
}

func NewSpideError(e error,msg string,url string) IVlinkError {
	er := SpideError{
		Url:url,
	}
	er.VlinkError=&VlinkError{
		BaseError: error2.NewBaseError(e, constants.SPIDE_ERROR, "爬虫爬取错误:"+msg),
	}
	return er
}
