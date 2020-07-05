/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 14:16 
# @File : base_service.go
# @Description : 
# @Attention : 
*/
package service

import (
	"fmt"
	"myLibrary/go-libary/go/log"
	"myLibrary/go-libary/go/utils"
	"runtime"
	"strings"
)

type BaseServiceImpl struct {
	MethodName string
	log        log.Logger
}

func NewBaseServiceImplWithLog4goLogger() *BaseServiceImpl {
	b := new(BaseServiceImpl)
	b.log = log.NewLog4goLogger(log.NewCommonBaseLoggerWithLog4go(utils.GenerateUUID()))
	return b
}

func (b *BaseServiceImpl) Info(first interface{}, info ...interface{}) {
	b.log.Info(first, info...)
}

func (b *BaseServiceImpl) Debug(first interface{}, info ...interface{}) {
	b.log.Debug(first, info...)
}

func (b *BaseServiceImpl) Error(first interface{}, info ...interface{}) {
	b.log.Error(first, info...)
}

func (b *BaseServiceImpl) SetPrefix(p string) {
	b.log.SetPrefix(p)
}

func (b *BaseServiceImpl) GetPrefix() string {
	return b.log.GetPrefix()
}

func (b *BaseServiceImpl) SetReqID(r string) {
	b.log.SetReqID(r)
}

func (b *BaseServiceImpl) GetReqID() string {
	return b.log.GetReqID()
}

func (b *BaseServiceImpl) BeforeStart(method string) {
	b.MethodName = method
	methodName := b.GetPrefix() + "->" + method
	b.log.SetPrefix(methodName)
	fmt.Println(strings.Repeat(">", 20))
	fmt.Println("开始调用:" + methodName)
	// b.Info("开始调用:" + methodName)
	fmt.Println(strings.Repeat(">", 20))
}

func (b *BaseServiceImpl) AfterEnd() {
	if err := recover(); err != nil {
		pc, _, lineNO, ok := runtime.Caller(1)

		if ok {
			b.Error("结束方法时, (%s:%d)出现panic:%s", runtime.FuncForPC(pc).Name(), lineNO, err)
		} else {
			b.Error("结束方法时,出现panic:%s", err)
		}
	}

	pre := b.log.GetPrefix()
	b.SetPrefix(strings.TrimRight(pre, " -> "+b.MethodName))
	fmt.Println(strings.Repeat("<", 20))
	// b.Info("结束对: { " + b.MethodName + " } 方法的调用")
	fmt.Println("结束对: { " + b.MethodName + " } 方法的调用")
	fmt.Println(strings.Repeat("<", 20))
}
