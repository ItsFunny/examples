/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/25 8:52 下午
# @File : lt_34_在排序数组中查找第一次和最后一次出现的元素.go
# @Description :
# @Attention :
*/
package offer

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	r := []int{-1, -1}
	r[0] = searchRangeHalf(nums, target, true)
	r[1] = searchRangeHalf(nums, target, false)
	return r
}
func searchRangeHalf(nums []int, target int, leftDirection bool) int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// 相等则继续往左走
			index = mid
			if leftDirection {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	// if nums[left] == target {
	// 	return left
	// }
	// if nums[right] == target {
	// 	return right
	// }
	return index
}
