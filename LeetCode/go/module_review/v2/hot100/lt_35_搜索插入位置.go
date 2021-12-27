/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/26 11:57 上午
# @File : lt_35_搜索插入位置.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 题目需要转换为:「在一个有序数组中找第一个大于等于 target的下标 (相当于是找第一个出现的位置)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if target < nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid
		} else {
			// 一直往左边找,因为是第一个出现的位置
			right = mid
		}
	}
	// 如果是基于 for left+1<right 的模板, 最后都是需要判断left,right的
	if nums[left] >= target {
		return left
	}
	if nums[right] >= target {
		return right
	} else {
		// 说明当前right达到了最大值,并且是整个数组中都没有这个数
		return right + 1
	}
}
