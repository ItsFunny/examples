/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-15 11:51 
# @File : base.go
# @Description : 
# @Attention : 
*/
package base

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"myLibrary/go-libary/go/converters"
	"vlink.com/v2/vlink-common/constants"
	error2 "vlink.com/v2/vlink-common/error"
)

type TransBaseType int

type MethodName string

// 代表区块链上的key
type Key string

// fromWalletAddress 从哪个钱包过来的
type From string

// toWalletAddress 去往哪个钱包的
type To string

// token 交易coin
type Token float64

type Version uint64

// 用于生成key的常量定义,如 userkey_constats 传入参数则会自动生成compositekey或者是单个key
type ObjectType string

type KeyGenerater func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string, error2.IVlinkError)

type IVlinkTxBaseResper interface {
	// 用于后续的钩子记录
	GetTXRecordInfo() TXRecordInfo
	// 用于获取内部数据
	GetReturnData() interface{}

	// 用于获取返回code,判断是否成功
	GetCode() int
	GetMsg() string
}

// version,from
type VlinkPeerResponse struct {
	peer.Response
	// Version             uint64
	// From                From
	// To                  To
	// Token               Token
	// BaseTransactionType TransBaseType
}

// 区块链调用的返回值
type BaseFabricResp struct {
	// 具体的业务返回值
	DataBytes []byte `json:"dataBytes"`
	// 代表code
	CodeBytes []byte `json:"codeBytes"`
	// 代表消息 ,直接String强转即可
	MsgBytes []byte `json:"msgBytes"`
	// 遗留字节,可能有其他用处
	OtherBytes []byte `json:"otherBytes"`
}

func NewSuccessBaseFabricResp() *BaseFabricResp {
	r := new(BaseFabricResp)
	r.CodeBytes = converter.BigEndianInt642Bytes(constants.SUCCESS)
	r.MsgBytes = []byte("SUCCESS")
	return r
}

func NewBaseFabricResp(code int, msg string) *BaseFabricResp {
	r := new(BaseFabricResp)
	r.CodeBytes = converter.BigEndianInt642Bytes(int64(code))
	r.MsgBytes = []byte(msg)
	return r
}

func Success(bytes []byte) VlinkPeerResponse {
	return SuccessWithDetail(bytes, constants.SUCCESS, "SUCCESS")
}
func Fail(e error2.IVlinkError) VlinkPeerResponse {
	resp := NewBaseFabricResp(int(e.GetCode()), e.GetMsg())
	bbs, _ := json.Marshal(resp)
	r := VlinkPeerResponse{
		Response: shim.Success(bbs),
	}
	return r
}
func SuccessWithDetail(bytes []byte, code int, msg string) VlinkPeerResponse {
	resp := NewBaseFabricResp(code, msg)
	resp.DataBytes = bytes
	bbs, _ := json.Marshal(resp)
	r := VlinkPeerResponse{
		Response: shim.Success(bbs),
	}
	return r
}

type BaseFabricAfterValidModel struct {
	Req           interface{}
	Version       uint64
	BaseTransType TransBaseType
}

type BCBaseNodeInfo struct {
	From       From          `json:"from"`
	To         To            `json:"to"`
	Token      Token         `json:"token"`
	Version    Version       `json:"version"`
	TxBaseType TransBaseType `json:"txBaseType"`
	// 是否是加密数据
	Encrypted bool `json:"encrypted"`
}

// 用于当业务结束之后,钩子对交易记录,既RecordTx
type TXRecordInfo struct {
	From  From  `json:"from"`
	To    To    `json:"to"`
	Token Token `json:"token"`
}

type BCBase struct {
	Key   Key   `json:"key"`
	From  From  `json:"from"`
	To    To    `json:"to"`
	Token Token `json:"token"`
}

type BCPutStateReq struct {
	BCBase
	// 是否需要加密
	NeedEncrypt bool
}

func (r BCPutStateReq) String() string {
	str := fmt.Sprintf("{ Key=[%+v] From=[%+v],To=[%+v],Token=[%+v],NeedEncrypt=[%+v] }", r.Key, r.From, r.To, r.To, r.NeedEncrypt)
	return str
}

type BCGetStateResp struct {
	BCBase
	// 是否需要解密,True代表update的时候需要加密
	NeedDecrypt bool
}

func GetRegularInfo(bs []byte) (BCBaseNodeInfo, []byte) {
	baseTypes := bs[0:8]
	fromWalletBytes := bs[8:42]
	toWalletBytes := bs[42:76]
	transAmount := bs[76:84]
	versionType := bs[84:92]

	m := BCBaseNodeInfo{
		From: From(fromWalletBytes),
		// To:         To(toWalletBytes),
		Token:      Token(converter.BigEndianBytesToFloat64(transAmount)),
		Version:    Version(converter.BigEndianBytes2Int64(versionType)),
		TxBaseType: TransBaseType(converter.BigEndianBytes2Int64(baseTypes)),
	}

	if toWalletBytes[0] == 0 && toWalletBytes[31] == 0 {
		m.To = ""
	} else {
		m.To = To(toWalletBytes)
	}
	if bs[92] == 1 {
		m.Encrypted = true
	}

	return m, bs[100:]
}
