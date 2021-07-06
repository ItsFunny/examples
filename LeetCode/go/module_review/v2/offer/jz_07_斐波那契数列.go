/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/29 9:01 上午
# @File : jz_07_斐波那契数列.go
# @Description :
# @Attention :
*/
package offer

func Fibonacci( n int ) int {
	// write code here
	if n==1{
		return 1
	}
	if n<=0{
		return 0
	}
	return Fibonacci(n-1)+Fibonacci(n-2)
}
