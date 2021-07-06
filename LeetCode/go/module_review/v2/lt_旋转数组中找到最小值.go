/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/23 9:20 上午
# @File : lt_旋转数组中找到最小值.go
# @Description :
# @Attention :
*/
package v2


func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)>>1
		if nums[mid] <= nums[end] {
			end=mid
		} else {
			start = mid
		}
	}
	if nums[end] < nums[start] {
		return nums[end]
	}
	return nums[start]
}
