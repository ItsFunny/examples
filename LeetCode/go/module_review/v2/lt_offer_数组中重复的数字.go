/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/25 9:00 上午
# @File : lt_offer_数组中重复的数字.go
# @Description :
# @Attention :
*/
package v2

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
