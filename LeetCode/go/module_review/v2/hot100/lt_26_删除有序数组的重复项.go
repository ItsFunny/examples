/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/17 9:04 上午
# @File : lt_26_删除有序数组的重复项.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 双指针:快慢指针
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow, fast := 0, 1
	for ; fast < len(nums); {
		if nums[fast] == nums[slow] {
			// 当重复的时候,快指针继续移动,直到遇到不相等的数
			fast++
		} else {
			slow++
			// 不重复的时候,这个元素移动到之前重复的元素位置处
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
