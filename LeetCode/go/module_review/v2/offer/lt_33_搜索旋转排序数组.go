/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/25 8:40 上午
# @File : lt_33_搜索旋转排序数组.go
# @Description :
# @Attention :
*/
package offer

import "math"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if target < nums[0] {
			// 说明target在右半段,则需要确保左半段都改为小值
			// 如果mid在左半段,组需要将mid移到右半段,所以把左
			if nums[mid] > nums[0] {
				nums[mid] = math.MinInt32
			}
		} else {
			// 说明target在左半段,则需要将mid 往左移,因此对于小于target[0]的数,都扩到最大
			if nums[mid] < nums[0] {
				nums[mid] = math.MaxInt32
			}
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}

func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}
