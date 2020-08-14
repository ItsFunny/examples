/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-14 09:02 
# @File : lt_61_Search_for_a_Range.go
# @Description : 
# @Attention : 
*/
package sort

func searchRange(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1
	result := []int{-1, -1}
	for start < end {
		mid := (start + end) >> 1
		if nums[mid] > target {
			start = mid
		} else if nums[mid] < target {
			end = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		result[0] = target
	} else if nums[end] == target {
		result[0] = target
	} else {
		return result
	}

	start = 0
	end = len(nums) - 1
	for start < end {
		mid := (start + end) >> 1
		if nums[mid] < target {
			end = mid
		} else if nums[mid] > target {
			start = mid
		} else {
			start = mid
		}
	}

	if nums[start] == target {
		result[1] = start
	} else if nums[end] == target {
		result[1] = end
	}

	return result
}
