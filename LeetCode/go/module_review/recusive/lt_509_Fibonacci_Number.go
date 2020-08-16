/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 17:13 
# @File : lt_509_Fibonacci_Number.go
# @Description : 
# @Attention : 
*/
package recusive

func fib(N int) int {
	return f(N)
}

func f(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return f(n-1) + f(n-2)
}
