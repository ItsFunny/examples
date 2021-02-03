/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:21 
# @File : log.go
# @Description : 
# @Attention : 
*/
package log


type Logger interface {
	Info(first interface{}, info ...interface{})
	Debug(first interface{}, info ...interface{})
	Error(first interface{}, info ...interface{})
	SetPrefix(p string)
	GetPrefix() string
	SetReqID(r string)
	GetReqID() string
}

type ConcreteLogger interface {
	RecordInfo(first interface{}, info ...interface{})
	RecordDebug(first interface{}, info ...interface{})
	RecordError(first interface{}, info ...interface{})
}


// type BaseLog struct {
// 	log Logger
// }
//
// func NewBaseLog(ll Logger) *BaseLog {
// 	l := new(BaseLog)
// 	l.log = ll
// 	return l
// }
//
// func (l *BaseLog) Info(first interface{}, info ...interface{}) {
// 	l.log.Info(first, info...)
// }
//
// func (l *BaseLog) Debug(first interface{}, info ...interface{}) {
// 	l.log.Debug(first, info...)
// }
//
// func (l *BaseLog) Error(first interface{}, info ...interface{}) {
// 	l.log.Error(first, info...)
// }
