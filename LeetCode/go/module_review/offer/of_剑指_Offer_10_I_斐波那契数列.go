/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 13:01 
# @File : of_剑指_Offer_10_I_斐波那契数列.go
# @Description : 
# @Attention : 
*/
package offer

func fib(n int) int {
	return f(n, map[int]int{})
}
func f(n int, m map[int]int) int {
	if n < 2 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	}
	prev := f(n-1, m) % 1000000007
	m[n-1] = prev
	pprev := f(n-2, m) % 1000000007
	m[n-2] = pprev
	res := (prev + pprev) % 1000000007
	m[n] = res
	return res
}
