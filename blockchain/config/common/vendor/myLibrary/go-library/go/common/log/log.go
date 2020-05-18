package log

import (
	"fmt"
	"runtime"

	"code.google.com/log4go"
)

type Log struct {
	InitLog
}

type InitLog struct {
	ReqID string

	prefix string
}

func NewLog(log InitLog) *Log {
	return &Log{InitLog: log}
}

func (l *Log) SetPrefix(str string) {
	l.prefix = str
}

func (l *Log) GetPrefix() string {
	return l.prefix
}

func (l Log) Info(first interface{}, info ...interface{}) {
	pc, _, lineNO, ok := runtime.Caller(1)
	src := ""

	if l.prefix != "" {
		src = fmt.Sprintf("[%s] ", l.prefix)
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
		log4go.Info(src, info...)
	} else {
		temp := []interface{}{first}
		info = append(temp, info...)
		log4go.Info(src, info)
	}
}

func (l Log) Debug(first interface{}, info ...interface{}) {
	pc, _, lineNO, ok := runtime.Caller(1)
	src := ""

	if l.prefix != "" {
		src = fmt.Sprintf("[%s] ", l.prefix)
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
		log4go.Debug(src, info...)
	} else {
		temp := []interface{}{first}
		info = append(temp, info...)
		log4go.Debug(src, info)
	}
}

func (l Log) Error(first interface{}, info ...interface{}) {
	pc, _, lineNO, ok := runtime.Caller(1)
	src := ""

	if l.prefix != "" {
		src = fmt.Sprintf("[%s] ", l.prefix)
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
		log4go.Error(src, info...)
	} else {
		temp := []interface{}{first}
		info = append(temp, info...)
		log4go.Error(src, info)
	}
}
