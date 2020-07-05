/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-02 09:48 
# @File : _45_Jump_Game_II.go
# @Description : 
# @Attention : 
*/
package main

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var maxIdx, i, step int
	for i < len(nums)-1 {
		nextIdx, tmpIdx := 0, 0
		for j := 0; j < nums[i]; j++ {
			tmpIdx = i + j + 1
			if tmpIdx < len(nums) && tmpIdx+nums[tmpIdx] > maxIdx {
				maxIdx = tmpIdx + nums[tmpIdx]
				nextIdx = tmpIdx
			}
			if tmpIdx == len(nums)-1 {
				nextIdx = tmpIdx
			}
		}

		i = nextIdx
		step++
	}

	return step
}