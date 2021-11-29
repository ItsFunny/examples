/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/20 10:45 上午
# @File : lt_169_多数元素.go
# @Description :
# @Attention :
*/
package offer

// 关键:
// 排序,找中间值
func majorityElement(nums []int) int {
	majorityElementQSort(nums, 0, len(nums)-1)
	return nums[len(nums)>>1]
}
func majorityElementQSort(nums []int, start, end int) {
	if start < end {
		paration := majorityElementQSortParation(nums, start, end)
		majorityElementQSort(nums, start, paration)
		majorityElementQSort(nums, paration+1, end)
	}
}

func majorityElementQSortParation(nums []int, start, end int) int {
	standard := nums[start]
	for start < end {
		for ; end > start && nums[end] >= standard; end-- {
		}
		nums[start] = nums[end]
		for ; start < end && nums[end] <= standard; start++ {
		}
		nums[end] = nums[start]
	}
	nums[start] = standard
	return start
}
