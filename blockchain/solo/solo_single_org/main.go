/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-25 16:07 
# @File : main.go
# @Description : 
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
}
