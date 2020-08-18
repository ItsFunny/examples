/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 13:15 
# @File : of_剑指_Offer_10_II_青蛙跳台阶问题.go
# @Description : 
# @Attention : 
*/
package offer

func numWays(n int) int {
	return ff(n, map[int]int{})
}
func ff(n int, m map[int]int) int {
	if n == 0 {
		return 1
	}
	if n < 3 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	}
	prev := ff(n-1, m) % 1000000007
	m[n-1] = prev
	pprev := ff(n-2, m) % 1000000007
	m[n-2] = pprev
	res := (prev + pprev) % 1000000007
	m[n] = res
	return res
}
