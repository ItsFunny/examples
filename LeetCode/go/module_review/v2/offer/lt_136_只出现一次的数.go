/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/21 8:57 上午
# @File : lt_136_只出现一次的数.go
# @Description :
# @Attention :
*/
package offer

func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}
