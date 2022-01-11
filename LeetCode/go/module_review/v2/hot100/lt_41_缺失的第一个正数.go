/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/6 9:13 上午
# @File : lt_41_缺失的第一个正数.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 将 1放在下标为0 的地方,2放在1的地方,类推
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] < len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums)+1
}

