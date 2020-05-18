/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-28 10:53 
# @File : error.go
# @Description :  error 封装
*/
package error

import "errors"

const (
	RESULT_LEVEL_SUCCESS = 1
	// 返回原生的错误,对原生的信息不包装成其他类型的错误
	RESULT_RETURN_NATIVE_ERR = RESULT_LEVEL_SUCCESS << 1
)

type WrapErrFunc func(int2 int, msg string) error

// TODO 通过map 实现对错误的自动转换, map[int]func(res BaseBussResult)error
var (
	errorWrapMap      map[int]WrapErrFunc
	ERR_NO_MATCH_WRAP = errors.New("无匹配的异常包装func")
)

func init() {
	AddWrap(RESULT_RETURN_NATIVE_ERR, func(int2 int, msg string) error {
		return errors.New(msg)
	})
}

func initData() {
	errorWrapMap = make(map[int]WrapErrFunc)
}
func AddWrap(level int, w WrapErrFunc) {
	if errorWrapMap == nil {
		initData()
	}
	errorWrapMap[level] = w
}
func GetLevelErr(level int, msg string) error {
	if f, exist := errorWrapMap[level]; exist {
		return f(level, msg)
	} else {
		return ERR_NO_MATCH_WRAP
	}
}

// 业务逻辑执行的结果
type BaseBussResult struct {
	// 执行结果,通过执行结果判断返回什么类型的error ,为
	ResultLevel int
}
