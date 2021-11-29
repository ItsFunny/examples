/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/30 9:13 上午
# @File : lt_42_接雨水.go
# @Description :
# @Attention :
*/
package offer

// 参考: https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode-solution-tuvc/
// 关键: 双指针
// 题目关键: 水的计算是通过,最小的来判断的
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	r := 0
	for left < right {
		leftMax = trapMax(leftMax, height[left])
		rightMax = trapMax(rightMax, height[right])
		if height[left] < height[right] {
			r += leftMax - height[left]
			left++
		} else {
			r += rightMax - height[right]
			right--
		}
	}
	return r
}
func trapMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
