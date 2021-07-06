/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/15 8:44 上午
# @File : lt_二进制_只出现一次的数.go
# @Description :
# @Attention :
*/
package v2

func singleNumber(nums []int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		r ^= nums[i]
	}
	return r
}
