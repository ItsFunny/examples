/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/22 8:48 上午
# @File : lt_二分_搜索插入位置.go
# @Description :
# @Attention :
*/
package v2

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)>>1
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] >= target {
		return start
	}
	if nums[end] >= target {
		return end
	} else if nums[end] < target {
		return end + 1
	}
	return 0
}
