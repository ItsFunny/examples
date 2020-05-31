/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-03-26 13:31 
# @File : main.go
# @Description : 
# @Attention : 
*/
package main

import (
	"examples/blockchain/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	yamlPath := "/Users/joker/go/src/examples/blockchain/twoorg_twochannel_chaincodes/application-dev.yaml"
	if e := config.Config(yamlPath); nil != e {
		panic(e)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	si := <-c
	if si != nil {
		fmt.Println("接收到结束信号")
		os.Exit(0)
	}
	fmt.Println("程序结束运行")
}
