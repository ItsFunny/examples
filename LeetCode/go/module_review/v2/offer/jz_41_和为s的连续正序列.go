/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/20 9:33 上午
# @File : jz_41_和为s的连续正序列.go
# @Description :
# @Attention :
*/
package offer

func FindContinuousSequence(sum int) [][]int {
	left, right := 1, 1
	tmp := 0
	r := make([][]int, 0)
	for left <= sum>>1 {
		if tmp < sum {
			tmp += right
			right++
		} else if tmp > sum {
			tmp -= left
			left++
		} else {
			v := make([]int, 0)
			for i := left; i < right; i++ {
				v = append(v, i)
			}
			r = append(r, v)
			tmp -= left
			left++
		}
	}
	return r
}
