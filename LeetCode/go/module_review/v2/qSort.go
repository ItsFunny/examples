/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/8 10:07 上午
# @File : qSort.go
# @Description :
# @Attention :
*/
package v2

func qSort(nums []int, start, end int) {
	if start < end {
		paration := qSortParation(nums, start, end)
		qSort(nums, start, paration)
		qSort(nums, paration, end)
	}
}
func qSortParation(nums []int, start, end int) int {
	standard := nums[start]
	for start < end {
		for ; start < end && nums[end] > standard; end-- {
		}
		nums[start] = nums[end]
		for ; start < end && nums[start] < standard; start++ {

		}
		nums[end] = nums[start]
	}
	nums[start] = standard
	return start
}
