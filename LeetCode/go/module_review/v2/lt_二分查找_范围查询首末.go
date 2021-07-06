/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/21 9:25 上午
# @File : lt_二分查找_范围查询首末.go
# @Description :
# @Attention :
*/
package v2

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)>>1
		if nums[mid] > target {
			start = mid
		} else if nums[mid] < target {
			end = mid
		} else {
			// 因为是第一次出现和最后一次出现,第一次肯定在左边,所以向左找
			end = mid
		}
	}
	r := []int{}
	if nums[start] == target {
		r = append(r, start)
	} else if nums[end] == target {
		r = append(r, end)
	} else {
		return []int{-1, -1}
	}

	start = 0
	end = len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)>>1
		if nums[mid] > target {
			start = mid
		} else if nums[mid] < target {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] == target {
		r = append(r, start)
	} else if nums[end] == target {
		r = append(r, end)
	} else {
		return []int{-1, -1}
	}
	return r
}
