/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/23 9:47 下午
# @File : lt_24_删除有序数组的重复项.go
# @Description :
# @Attention :
*/
package offer

// 关键
// 题目特点: 有序+重复
// 解题关键: 快慢指针,慢指针充当不重复的元素个数,快指针快速过滤
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow, fast := 0, 1
	for ; fast < len(nums); {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			slow++
			// 移动到下个匹配重复的地方,因为是有序的,下一个重复的必然>=当前值
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
