/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-14 08:59 
# @File : base.go
# @Description : 
# @Attention : 
*/
package sort

// 二分搜索常用模板

// 二分搜索最常用模板
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (start + end) >> 1
		if nums[mid] == target {
			return 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
