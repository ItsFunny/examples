/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-15 13:44 
# @File : lt_33_Search_in_Rotated_Sorted_Array.go
# @Description :

# @Attention :
	在每个有序子数组之间查找
*/
package half

func search(nums []int, target int) int {
	index := findReverseInde(nums)
	val:=0
	val=find(nums[:index],target)
	if val!=-1{
		return val
	}
	val=find(nums[index:],target)
	return val
}
func find(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)>>1
		if target < nums[mid] {
			end = mid
		} else if target == nums[mid] {
			return mid
		} else {
			start = mid
		}
	}
	return -1
}
func findReverseInde(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)>>1
		// 当大于的时候是不正常的
		if nums[mid] > nums[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if end < start {
		return end
	}
	return start
}
