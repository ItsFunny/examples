/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/20 9:09 上午
# @File : jz_40_数组中只出现一次的两个数字.go
# @Description :
# @Attention :
*/
package offer

func FindNumsAppearOnce(array []int) []int {
	r := make([]int, 2)
	mix := 0
	for _, v := range array {
		mix ^= v
	}
	r[0], r[1] = mix, mix
	mix = (mix & (mix - 1)) ^ mix
	for _, v := range array {
		if v&mix == 0 {
			r[0] ^= v
		} else {
			r[1] ^= v
		}
	}
	if r[0] > r[1] {
		r[1], r[0] = r[0], r[1]
	}
	return r
}
