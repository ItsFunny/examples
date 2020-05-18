/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-15 11:49 
# @File : args.go
# @Description : 
# @Attention : 
*/
package service

import error2 "vlink.com/v2/vlink-common/error"




type ArgsHelper interface {
	Checker(args []string) error2.IVlinkError
	Converter(args []string) (interface{}, error2.IVlinkError)
}
