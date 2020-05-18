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
	"myLibrary/go-library/go/authentication"
	"myLibrary/go-library/go/converters"
	"vlink.com/v2/vlink-common/constants"
	error2 "vlink.com/v2/vlink-common/error"
)

type TransBaseType int
type ChannelID string
type OrganizationID string
type ChainCodeID string

type ChainBaseReq struct {
	MethodName  MethodName
	ChannelID   ChannelID
	ChainCodeID ChainCodeID
}

type TransBaseDescription struct {
	TransBaseType TransBaseTypeV2
	Description   string
}

func NewNeedRecordTransBaseDescription(baseValue TransBaseTypeV2Value, desc string) TransBaseDescription {
	description := TransBaseDescription{
		Description: desc,
	}
	description.TransBaseType = CreateNeedRecordBaseType(baseValue)
	return description
}

func NewUnRecordTransBaseDescription(baseValue TransBaseTypeV2Value, desc string) TransBaseDescription {
	description := TransBaseDescription{
		Description: desc,
	}
	description.TransBaseType = CreateUnNeedRecordBaseType(baseValue)
	return description
}

func ConvBytes2TransBaseTypeV2(bytes []byte) TransBaseTypeV2 {
	authorities, _ := authentication.BigEndianConvtBytes2Authority(bytes)
	return TransBaseTypeV2(authorities)
}

// func NewTransBaseDescription(baseType TransBaseTypeV2Value, desc string) TransBaseDescription {
// 	description := TransBaseDescription{
// 		Description: desc,
// 	}
// 	description.TransBaseType = NewTransBaseTypeV2()
// 	return description
// }

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
	GetTXRecordInfoList() []TxRecordInfoNode
	// 用于获取内部数据
	GetReturnData() interface{}

	// 用于获取返回code,判断是否成功
	GetCode() int
	GetMsg() string

	// 2020-02-24
	// 15:37 update
	// 设置相关的 tx信息
	SetTxBaseType(baseType TransBaseTypeV2)
	SetTxDescription(d string)

	// 2020-03-16
	// 14:57 add
	SetChannelID(c ChannelID)
	SetTransactionID(tid string)
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
	// 日志字节,包含了是否需要记录等信息
	LogBytes []byte `json:"logBytes"`
}
type ResultInfo struct {
	// code 要么 与success 做了或运算代表成功,
	// 要么与 fail 做了或运算代表逻辑失败,至于是否显示抛出取决于是否与out_put做了或运算
	LogicCode  int    `json:"code"`
	LogicMsg   string `json:"msg"`
	From       From   `json:"from"`
	To         To     `json:"to"`
	Token      Token  `json:"token"`
	DataBytes  []byte `json:"data"`
	OtherBytes []byte `json:"otherBytes"`
	// 2020-03-16 13:16 add
	// 为了后续业务的变更,这里统一返回数组的形式返回,而非返回单个数据的形式,便于后续维护
	TxRecords []TxRecordInfoNode `json:"txRecords"`
}

func (r ResultInfo) String() string {
	str := "code=[%d],msg=[%s],from=[%v],to=[%v],token=[%v],data=[%s],other=[%s]"
	return fmt.Sprintf(str, r.LogicCode, r.LogicMsg, r.From, r.To, r.Token, string(r.DataBytes), string(r.OtherBytes))
}

func (r ResultInfo) Success() bool {
	return (r.LogicCode & constants.SUCCESS) > 0
}

// "from":"13Yrapjusm3ic9ncVGrCHXiNJPXJpwMQCD",
// "to":"",
// "token":0,
// "Data":{"dna":"840aa0eb-96f5-4960-914c-1b313570ccca","prvKey":"","coinAddress":"1MxWsr2qXr63JVfLghR3u5E3asJP7tutr9"},
// "Code":1,"Msg":"SUCCESS
func (r BaseFabricResp) Convt2ResultInfo() ResultInfo {
	res := ResultInfo{}
	res.LogicCode = int(converter.BigEndianBytes2Int64(r.CodeBytes))
	res.LogicMsg = string(r.MsgBytes)
	res.OtherBytes = r.OtherBytes
	res.DataBytes = r.DataBytes
	json.Unmarshal(r.LogBytes, &res.TxRecords)

	return res
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
	return SuccessWithDetail(bytes, nil, constants.SUCCESS, "SUCCESS")
}
func Fail(e error2.IVlinkError) VlinkPeerResponse {
	resp := NewBaseFabricResp(int(e.GetCode()), e.GetMsg())
	bbs, _ := json.Marshal(resp)
	r := VlinkPeerResponse{
		Response: shim.Success(bbs),
	}
	return r
}

func SuccessWithDetail(bytes []byte, logBytes []byte, code int, msg string) VlinkPeerResponse {
	resp := *NewBaseFabricResp(code, msg)
	resp.DataBytes = bytes
	resp.LogBytes = logBytes
	bbs, _ := json.Marshal(resp)
	r := VlinkPeerResponse{
		Response: shim.Success(bbs),
	}
	return r
}

type BaseFabricAfterValidModel struct {
	Req           interface{}
	Version       uint64
	BaseTransType TransBaseTypeV2
	// 描述
	BaseTransDescription string
}

type BCBaseNodeInfo struct {
	From       From            `json:"from"`
	To         To              `json:"to"`
	Token      Token           `json:"token"`
	Version    Version         `json:"version"`
	TxBaseType TransBaseTypeV2 `json:"txBaseType"`
	// 是否是加密数据
	Encrypted bool `json:"encrypted"`
}

// 用于当业务结束之后,钩子对交易记录,既RecordTx
type TXRecordInfo struct {
	From          From            `json:"from"`
	To            To              `json:"to"`
	Token         Token           `json:"token"`
	BaseType      TransBaseTypeV2 `json:"baseType"`
	TxDescription string          `json:"txDescription"`
	// 在哪条链上
	ChannelID ChannelID `json:"channelId"`
	// 2020-02-27
	// 10:23 add
	// 交易ID
	TransactionID string `json:"transactionId"`
}

type TxRecordInfoNode struct {
	From          From            `json:"from"`
	To            To              `json:"to"`
	Token         Token           `json:"token"`
	BaseType      TransBaseTypeV2 `json:"baseType"`
	TxDescription string          `json:"txDescription"`
	// 在哪条链上
	ChannelID ChannelID `json:"channelId"`
	// 2020-02-27
	// 10:23 add
	// 交易ID
	TransactionID string `json:"transactionId"`
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

// 废弃了
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
		TxBaseType: ConvBytes2TransBaseTypeV2(baseTypes),
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

func GetRegularInfoV2(bs []byte) (BCBaseNodeInfo, []byte) {
	defer func() {
		recover()
	}()
	// baseTypes := bs[0:8]
	l := len(bs)
	fromWalletBytes := bs[constants.FROM_WALLET_ADDRESS_BEGIN:constants.FROM_WALLET_ADDRESS_EDN]
	toWalletBytes := bs[constants.TO_WALLET_ADDRESS_BEGIN:constants.TO_WALLET_ADDRESS_END]
	transAmount := bs[constants.TRANS_AMOUNT_INDEX_BEGIN:constants.TRANS_AMOUNT_INDEX_END]
	versionType := bs[constants.VERSION_TYPE_INDEX_BEGIN:constants.VERSION_TYPE_INDEX_END]
	baseTypeLength := bs[constants.BASE_TYPE_BYTE_INDEX]

	typeIndex := l - int(baseTypeLength*8)
	baseTypes := bs[typeIndex:l]

	m := BCBaseNodeInfo{
		From: From(fromWalletBytes),
		// To:         To(toWalletBytes),
		Token:      Token(converter.BigEndianBytesToFloat64(transAmount)),
		Version:    Version(converter.BigEndianBytes2Int64(versionType)),
		TxBaseType: ConvBytes2TransBaseTypeV2(baseTypes),
	}

	if toWalletBytes[0] == 0 && toWalletBytes[31] == 0 {
		m.To = ""
	} else {
		m.To = To(toWalletBytes)
	}
	if bs[constants.CRYPT_INDEX] == 1 {
		m.Encrypted = true
	}

	return m, bs[constants.VLINK_COMMON_INDEX_END:typeIndex]
}

type InvokeBaseReq struct {
	ChannelName   string
	ChaincodeName string
	MethodName    MethodName
	Data          interface{}
}
