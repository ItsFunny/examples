/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-12 10:39 
# @File : _15_3Sum.go
# @Description : Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
三数求和为0
# @Attention : 
*/
package main

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// for why  ?
		if i > 0 && nums[i] == nums[i-1] {continue}
		j, k := i + 1, len(nums) - 1
		for j < k {
			if j > i + 1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums) - 1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}