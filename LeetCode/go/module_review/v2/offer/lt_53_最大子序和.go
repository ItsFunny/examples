/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/1 8:43 上午
# @File : lt_53_最大子序和.go
# @Description :
# @Attention :
*/
package offer

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for index := 1; index < len(nums); index++ {
		if nums[index]+nums[index-1] > nums[index] {
			nums[index] = nums[index] + nums[index-1]
		}
		if nums[index] > max {
			max = nums[index]
		}
	}
	return max
}
