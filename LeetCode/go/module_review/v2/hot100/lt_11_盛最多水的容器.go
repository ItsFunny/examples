/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/11 8:43 上午
# @File : lt_11_盛最多水的容器.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 计算公式: 面积=长*宽 ,在宽移动的时候必然缩小的前提下,则往高处移动,才有可能计算出更大的面积
// 左右双指针
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ret := 0

	for left < right {
		// 并且注意: 水的计算,是要根据木桶的最短板的,所有 后面的被乘的数 是一个min的值
		ret = maxAreaMax(ret, (right-left)*maxAreaMin(height[left], height[right]))
		// 宽度注定减少的情况下,尽量往高的移动
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ret
}
func maxAreaMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxAreaMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
