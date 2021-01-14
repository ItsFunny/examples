/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 14:52 
# @File : of_剑指_Offer_21_调整数组顺序使奇数位于偶数前面.go
# @Description : 双指针
# @Attention : 
*/
package offer

func exchange(nums []int) []int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast]&1 == 1 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}
