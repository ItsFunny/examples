/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-27 07:34 
# @File : enroll.go
# @Description :    注册用户
# @Attention : 
*/
package test

import (
	"examples/blockchain/solo/solo_single_org/config"
	"flag"
	"fmt"
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
	// Enroll()

}

func Enroll(username,pwd string) {
	if enrolle := config.Enrolle(username,pwd); nil != enrolle {
		panic(enrolle)
	}else{
		fmt.Println("enroll 成功")
	}
}
