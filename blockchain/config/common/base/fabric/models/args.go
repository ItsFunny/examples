/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-15 11:57 
# @File : args.go
# @Description : 
# @Attention : 
*/
package models

import error2 "vlink.com/v2/vlink-common/error"

type ArgsChecker = func(args []string) error2.IVlinkError
type ArgsConverter = func(args []string) (interface{}, error2.IVlinkError)
type ArgsDecrypter = func(data interface{}, version string) (interface{}, error2.IVlinkError)

type ArgsParameter struct {
	ArgsChecker   ArgsChecker
	ArgsConverter ArgsConverter
}

var (
	DefaultNumberChecker = func(args []string) error2.IVlinkError {
		if len(args) < 1 {
			return error2.NewArguError(nil, "参数长度必须大于1")
		}
		return nil
	}
)

func NewDefaultCheckerParameter(ArgsConverter ArgsConverter) ArgsParameter {
	a := ArgsParameter{}
	a.ArgsChecker = DefaultNumberChecker
	a.ArgsConverter=ArgsConverter

	return a

}
