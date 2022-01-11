/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/10 9:29 上午
# @File : lt_42_接雨水.go
# @Description :
# @Attention :
*/
package hot100

// 关键:
// 1. 双指针
// 2. 面积= 长* 宽, 在宽注定减少的情况下,往 高的移动
func trap(height []int) int {
	ret := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = trapMax(leftMax, height[left])
		rightMax = trapMax(rightMax, height[right])
		// 注意这一步: ,匹配的是max,可以认为是使得得出,左右两边有低谷
		if leftMax < rightMax {
			ret += leftMax - height[left]
			left++
		} else {
			ret += rightMax - height[right]
			right++
		}
	}
	return ret
}

func trapMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
