/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-19 16:58 
# @File : gc2.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("end")
}
