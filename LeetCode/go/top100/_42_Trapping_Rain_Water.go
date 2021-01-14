/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-25 11:43 
# @File : _42_Trapping_Rain_Water.go
# @Description : 
# @Attention : 
*/
package main

func trap(height []int) int {
	left, maxLeft, right, maxRight := 0, 0, len(height)-1, 0
	res := 0
	for left < right {
		// 说明计算左边
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			// 说明计算右边
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}

	return res
}
