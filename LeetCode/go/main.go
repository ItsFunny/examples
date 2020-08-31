/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-09 09:59 
# @File : main.go
# @Description : 
# @Attention : 
*/
package main

import "fmt"

func main() {
	c := make(chan int, 10)
	c <- 1
	fmt.Println(<-c)
}
