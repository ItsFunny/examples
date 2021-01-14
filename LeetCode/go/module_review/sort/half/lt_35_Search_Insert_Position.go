/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-14 09:21 
# @File : lt_35_Search_Insert_Position.go
# @Description : 
# @Attention : 
*/
package half

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
