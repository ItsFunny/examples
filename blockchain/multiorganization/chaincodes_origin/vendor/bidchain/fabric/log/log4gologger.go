/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 12:33
# @File : log4gologger.go
# @Description :
# @Attention :
*/
package log

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"runtime"
)

type log4goLogger struct {
	*CommonBaseLogger
}

func NewLog4goLogger(b *CommonBaseLogger) *log4goLogger {
	l := new(log4goLogger)
	l.CommonBaseLogger = b
	return l
}

func (l *log4goLogger) RecordInfo(first interface{}, info ...interface{}) {
	pc, _, lineNO, ok := runtime.Caller(3)

	// _, _, _, ok := runtime.Caller(1)
	src := ""

	if l.Prefix != "" {
		src = fmt.Sprintf("[%s] ", l.Prefix)
	}

	if ok {
		src += fmt.Sprintf("[requestId:%v] (%s:%d)", l.ReqID, runtime.FuncForPC(pc).Name(), lineNO)
		// src += fmt.Sprintf("[requestId:%v]", l.ReqID)
	}

	isStr := false
	switch temp := first.(type) {
	case string:
		isStr = true
		src += temp
	}

	if isStr {
		// if len(info) > 0 {
		// 	msg := fmt.Sprint(fmt.Sprintf(strings.Repeat(" %v", len(info)), info...))
		// 	src += msg
		// }
		// logrus.Info(src)
		log4go.Info(src, info...)
	} else {
		temp := []interface{}{first}
		info = append(temp, info...)
		log4go.Info(src, info)
	}
}

func (l *log4goLogger) RecordDebug(first interface{}, info ...interface{}) {
	pc, _, lineNO, ok := runtime.Caller(1)
	src := ""

	if l.Prefix != "" {
		src = fmt.Sprintf("[%s] ", l.Prefix)
	}

	if ok {
		src += fmt.Sprintf("[requestId:%v] (%s:%d)", l.ReqID, runtime.FuncForPC(pc).Name(), lineNO)
	}

	isStr := false
	switch temp := first.(type) {
	case string:
		isStr = true
		src += temp
	}

	if isStr {
		// if len(info) > 0 {
		// 	msg := fmt.Sprint(fmt.Sprintf(strings.Repeat(" %v", len(info)), info...))
		// 	src += msg
		// }
		// logrus.Debug(src)
		log4go.Debug(src,info...)
	} else {
		temp := []interface{}{first}
		info = append(temp, info...)
		log4go.Debug(src, info)
	}
}

func (l *log4goLogger) RecordError(first interface{}, info ...interface{}) {
	pc, _, lineNO, ok := runtime.Caller(1)
	src := ""

	if l.Prefix != "" {
		src = fmt.Sprintf("[%s] ", l.Prefix)
	}

	if ok {
		src += fmt.Sprintf("[requestId:%v] (%s:%d)", l.ReqID, runtime.FuncForPC(pc).Name(), lineNO)
	}

	isStr := false
	switch temp := first.(type) {
	case string:
		isStr = true
		src += temp
	}

	if isStr {
		// if len(info) > 0 {
		// 	msg := fmt.Sprint(fmt.Sprintf(strings.Repeat(" %v", len(info)), info...))
		// 	src += msg
		// }
		// logrus.Error(src)
		log4go.Error(src,info...)
	} else {
		temp := []interface{}{first}
		info = append(temp, info...)
		log4go.Error(src, info)
	}
}
