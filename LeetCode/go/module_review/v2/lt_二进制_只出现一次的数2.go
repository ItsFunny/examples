/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/15 9:02 上午
# @File : lt_二进制_只出现一次的数2.go
# @Description :
# @Attention :
*/
package v2

func singleNumber2(nums []int) int {
	r := 0
	for i := 0; i < 64; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j]>>i&1 > 0 {
				count++
			}
		}
		r |= (count % 3) << i
	}

	return r
}
