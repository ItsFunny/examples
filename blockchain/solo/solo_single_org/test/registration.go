/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-05 13:18 
# @File : registration.go
# @Description : 
# @Attention : 
*/
package test

import (
	"examples/blockchain/solo/solo_single_org/config"
	"flag"
	"fmt"
)

func TestRegistration() {
	flag.Parse()
	e := config.Config("/Users/joker/go/src/examples/blockchain/solo/solo_single_org/config/application.yaml")
	if nil != e {
		panic(e)
	}
	Enroll("admin", "adminpw")
	Registration("joker", "123456")

	Enroll("joker","123456")
	config.NewAndInvoke()

}

func Registration(userName, pwd string) {
	if enrolle := config.Registraiton(userName, pwd); nil != enrolle {
		panic(enrolle)
	} else {
		fmt.Println("register 成功")
	}
}
