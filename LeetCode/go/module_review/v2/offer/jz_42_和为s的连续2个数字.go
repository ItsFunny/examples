/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/21 9:11 上午
# @File : jz_42_和为s的连续2个数字.go
# @Description :
# @Attention :
*/
package offer

import "math"

func FindNumbersWithSum(array []int, sum int) []int {
	left, right := 0, len(array)-1

	var v int
	min := math.MaxInt32
	var r []int
	for left < right {
		v = array[left] + array[right]
		if v < sum {
			left++
		} else if v > sum {
			right--
		} else {
			if (array[left] * array[right]) < min {
				r = []int{array[left], array[right]}
				min = array[left] * array[right]
			}
			left++
		}
	}
	return r
}
