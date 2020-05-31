/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-17 15:40 
# @File : resp.go
# @Description : 
# @Attention : 
*/
package models

import (
	"examples/blockchain/config/common/base/fabric"
	"examples/blockchain/config/common/constants"
)

type ILogicBaseService interface {
	GetLogicCode() int
	SetLogicCode(code int)

	GetFrom() base.From
	SetFrom(f base.From)

	GetToken() base.Token
	SetToken(t base.Token)

	GetTo() base.To
	SetTo(t base.To)

	GetLogicMsc() string
	SetLogcMsg(msg string)
}

type BaseResp struct {
	LogicCode int    `json:"code"`
	LogicMsg  string `json:"msg"`
}

// 用于底层的blockchain的统一内部base 返回值
// FIXME 更改为切片的形式
type LogicBaseResp struct {
	// 判断是否是业务异常,通常需要外抛出去
	LogicCode int       `json:"code"`
	LogicMsg  string    `json:"msg"`
	// From      base.From `json:"from"`
	// To        base.To
	// Token     base.Token
	// 2020-03-16 12:54 add
	// 业务需求还是需要将交易转换为数组的形式,而非单行记录
	TxRecords []base.TxRecordInfoNode `json:"txRecords"`
}

func (r LogicBaseResp) Success() bool {
	return (r.LogicCode & constants.SUCCESS) > 0
}


type LogicBaseRespNode struct {
	// 判断是否是业务异常,通常需要外抛出去
	LogicCode int       `json:"code"`
	LogicMsg  string    `json:"msg"`
	From      base.From `json:"from"`
	To        base.To
	Token     base.Token
}

// 2020-02-24 14:10 add
// 用于rpc 的base 返回值
type RPCLogicBaseResp struct {
	LogicCode int    `json:"code"`
	LogicMsg  string `json:"msg"`
	LogBytes  string `json:"logBytes"`

	// 2020-02-27
	// 10:16 add 交易id
	TransactionID string `json:"transactionId"`
}

// 2020-02-24 14:03 add
// 用于底层的logicService的返回值
type ServiceLogicBaseResp struct {
	LogicCode int    `json:"code"`
	LogicMsg  string `json:"msg"`
	LogBytes  []byte `json:"logBytes"`
}

func (l *LogicBaseResp) SetLogicCode(code int) {
	l.LogicCode = code
}

func (l *LogicBaseResp) SetFrom(f base.From) {
	panic("")
}

func (l *LogicBaseResp) SetToken(t base.Token) {
	panic("")
}

func (l *LogicBaseResp) SetTo(t base.To) {
	panic("")
}

func (l *LogicBaseResp) SetLogcMsg(msg string) {
	l.LogicMsg = msg
}

func (l LogicBaseResp) GetFrom() base.From {
	panic("")
}

func (l LogicBaseResp) GetToken() base.Token {
	panic("")
}

func (l LogicBaseResp) GetTo() base.To {
	panic("")
}

func NewLogicBaseResp() LogicBaseResp {
	l := LogicBaseResp{}
	l.LogicCode = constants.SUCCESS
	l.LogicMsg = "SUCCESS"
	l.TxRecords = make([]base.TxRecordInfoNode, 0)
	l.TxRecords=append(l.TxRecords,base.TxRecordInfoNode{})
	// l.LogicCode = code
	// l.LogicMsg = msg
	return l
}
func NewLogicFailBaseResp(msg string)LogicBaseResp{
	l := LogicBaseResp{}
	l.LogicCode = constants.FAIL
	l.LogicMsg=msg;
	return l
}
func (this *LogicBaseResp) SetOwnTxFrom(from base.From) {
	this.TxRecords[0].From = from
}

func (this *LogicBaseResp) SetOwnTo(to base.To) {
	this.TxRecords[0].To = to
}

func (this *LogicBaseResp) SetOwnToken(token base.Token) {
	this.TxRecords[0].Token = token
}

// 设置本业务的信息
func (this LogicBaseResp) SetOwnTxInfo(From base.From, to base.To, token base.Token) {
	this.TxRecords[0].From = From
	this.TxRecords[0].To = to
	this.TxRecords[0].Token = token
}

func NewRPCLogicBaseResp() *RPCLogicBaseResp {
	l := new(RPCLogicBaseResp)
	return l
}
func NewSuccessLogicBaseResp() *LogicBaseResp {
	l := new(LogicBaseResp)
	l.LogicCode = constants.SUCCESS
	l.LogicMsg = "SUCCESS"

	return l
}

func (l LogicBaseResp) GetLogicCode() int {
	return l.LogicCode
}

func (l LogicBaseResp) GetLogicMsc() string {
	return l.LogicMsg
}

// 2020-01-13 update ,以后统一用该返回值来代替交互
type InvokeResp struct {
	BaseResp
	// 返回值
	Data interface{}
}
