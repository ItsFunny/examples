/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/2 8:47 上午
# @File : lt_55_跳跃游戏.go
# @Description :
# @Attention :
*/
package offer

// 解题关键: 看能跳到最远的格子能否大于当前长度
func canJump(nums []int) bool {
	max := 0
	for index := range nums {
		if index > max {
			return false
		}
		// 之所以是nums[index]+index 而不是nums[index]+max 是因为,计算的是当前格子的长度,而不是总的
		if nums[index]+index > max {
			max = nums[index] + index
		}
	}
	return true
}
