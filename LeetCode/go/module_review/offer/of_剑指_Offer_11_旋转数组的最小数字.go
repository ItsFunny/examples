/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 13:29 
# @File : of_剑指_Offer_11_旋转数组的最小数字.go
# @Description : 
# @Attention : 
*/
package offer

func minArray(numbers []int) int {
	start := 0
	end := len(numbers) - 1
	for start < end {
		mid := start + (end-start)>>1
		if numbers[mid] < numbers[end] {
			end = mid
		} else if numbers[mid] > numbers[end] {
			start = mid + 1
		} else {
			end--
		}
	}
	return numbers[start]
}
