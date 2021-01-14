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
	"vlink.com/v2/vlink-common/base/fabric"
	"vlink.com/v2/vlink-common/constants"
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

type LogicBaseResp struct {
	// 判断是否是业务异常,通常需要外抛出去
	LogicCode int       `json:"code"`
	LogicMsg  string    `json:"msg"`
	From      base.From `json:"from"`
	To        base.To
	Token     base.Token
}

func (l *LogicBaseResp) SetLogicCode(code int) {
	l.LogicCode = code
}

func (l *LogicBaseResp) SetFrom(f base.From) {
	l.From = f
}

func (l *LogicBaseResp) SetToken(t base.Token) {
	l.Token = t
}

func (l *LogicBaseResp) SetTo(t base.To) {
	l.To = t
}

func (l *LogicBaseResp) SetLogcMsg(msg string) {
	l.LogicMsg = msg
}

func (l LogicBaseResp) GetFrom() base.From {
	return l.From
}

func (l LogicBaseResp) GetToken() base.Token {
	return l.Token
}

func (l LogicBaseResp) GetTo() base.To {
	return l.To
}

func NewLogicBaseResp() *LogicBaseResp {
	l := new(LogicBaseResp)
	// l.LogicCode = code
	// l.LogicMsg = msg
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

