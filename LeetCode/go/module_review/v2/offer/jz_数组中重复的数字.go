/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/27 9:06 上午
# @File : jz_数组中重复的数字.go
# @Description :
# @Attention :
*/
package offer

// 2,1,3,1,4

func duplicate(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for numbers[i] != i {
			if numbers[numbers[i]] == numbers[i] {
				return numbers[i]
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}
	return -1
}
