/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/30 9:12 上午
# @File : jz_11_二进制中1的个数.go
# @Description :
# @Attention :
*/
package offer

func NumberOf1(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
	}
	return count
}
