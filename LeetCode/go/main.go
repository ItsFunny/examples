/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-09 09:59 
# @File : main.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
)

func main() {
	// fmt.Println((-1 * math.MaxInt32) + (-1 * math.MaxInt32))

	a := 12
	fmt.Println(a)
	fmt.Println(a & (a - 1))
	fmt.Println((a & (a - 1)) ^ a)
}
