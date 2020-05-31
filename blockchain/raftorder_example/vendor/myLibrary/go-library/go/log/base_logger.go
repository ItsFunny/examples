/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:24 
# @File : logger.go
# @Description : 
# @Attention : 
*/
package log


type CommonBaseLogger struct {
	ReqID          string
	Prefix         string
	ConcreteLogger ConcreteLogger
}

// func NewCommonBaseLogger(reqID string) *CommonBaseLogger {
// 	l := new(CommonBaseLogger)
//
// 	goLogger := NewLog4goLogger(l)
// 	l.ConcreteLogger = goLogger
// 	l.ReqID = reqID
//
// 	return l
// }

func NewCommonBaseLoggerWithLog4go(reqID string)*CommonBaseLogger{
	l := new(CommonBaseLogger)

	goLogger := NewLog4goLogger(l)
	l.ConcreteLogger = goLogger
	l.ReqID = reqID

	return l
}

func (l *CommonBaseLogger) Info(first interface{}, info ...interface{}) {
	l.ConcreteLogger.RecordInfo(false, info...)
}

func (l *CommonBaseLogger) Debug(first interface{}, info ...interface{}) {
	l.ConcreteLogger.RecordDebug(first, info...)
}

func (l *CommonBaseLogger) Error(first interface{}, info ...interface{}) {
	l.ConcreteLogger.RecordError(first, info...)
}

func (l *CommonBaseLogger) SetPrefix(p string) {
	l.Prefix = p
}

func (l *CommonBaseLogger) GetPrefix() string {
	return l.Prefix
}

func (l *CommonBaseLogger) SetReqID(r string) {
	l.ReqID = r
}

func (l *CommonBaseLogger) GetReqID() string {
	return l.ReqID
}
