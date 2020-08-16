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
	// fmt.Println((-1 * math.MaxInt32) + (-1 * math.MaxInt32))
	m := make(map[string]int, 0)
	m["1"] = 1
	delete(m, "1")
	fmt.Println("delete")
}
