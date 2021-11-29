/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/26 9:31 上午
# @File : lt_217_存在重复的元素.go
# @Description :
# @Attention :
*/
package v2

// 关键:
// 可以用hashSet来做,也可以排序然后比较前后来做
func containsDuplicate(nums []int) bool {
	containsDuplicateQSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
func containsDuplicateQSort(nums []int, start, end int) {
	if start < end {
		paration := containsDuplicateQSortParation(nums, start, end)
		containsDuplicateQSort(nums, start, paration)
		containsDuplicateQSort(nums, paration+1, end)
	}
}
func containsDuplicateQSortParation(nums []int, start, end int) int {
	standard := nums[start]
	for start < end {
		for ; end >= start && nums[end] > standard; {
			end--
		}
		nums[start] = nums[end]
		for ; start <= end && nums[start] < standard;  {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = standard
	return start
}
