/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/4 9:43 下午
# @File : lt_66_加一.go
# @Description :
# @Attention :
*/
package offer

func plusOne(digits []int) []int {
	l := len(digits) - 1
	for i := l; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	ret := make([]int, len(digits)+1)
	ret[0] = 1

	return ret
}
