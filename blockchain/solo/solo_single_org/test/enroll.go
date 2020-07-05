/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-27 07:34 
# @File : enroll.go
# @Description :    注册用户
# @Attention : 
*/
package main

import (
	"examples/blockchain/solo/solo_single_org/config"
	"flag"
)

var (
	configPath = flag.String("config", "/Users/joker/go/src/examples/blockchain/solo/solo_single_org/config/application.yaml", "配置路径")
)

func main() {
	flag.Parse()
	e := config.Config(*configPath)
	if nil != e {
		panic(e)
	}
	enroll()

}

func enroll() {
	if enrolle := config.Enrolle("joker", "123"); nil != enrolle {
		panic(enrolle)
	}
}
