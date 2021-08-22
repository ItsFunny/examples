/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/13 9:11 上午
# @File : jz_31_整数中1出现的次数.go
# @Description :
# @Attention :
*/
package offer

func NumberOf1Between1AndN_Solution(n int) int {
	r := 0
	for i := 0; i <= n; i++ {
		if i&1 >= 1 {
			r++
		}
	}
	return r
}
