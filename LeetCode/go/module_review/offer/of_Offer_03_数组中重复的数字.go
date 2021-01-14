/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-13 09:40 
# @File : of_Offer_03_数组中重复的数字.go
# @Description :
	在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
	数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
# @Attention :
*/
package offer

func findRepeatNumber(nums []int) int {
	array := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if array[nums[i]]> 0 {
			return nums[i]
		}
		array[nums[i]]++
	}
	return -1
}
