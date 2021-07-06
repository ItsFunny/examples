/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/16 10:04 上午
# @File : lt_二进制_位1的个数.go
# @Description :
# @Attention :
*/
package v2

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count++
		num = num & (num - 1)
	}
	return count
}
