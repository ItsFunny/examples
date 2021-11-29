/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/17 9:55 下午
# @File : lt_11_盛最多水的容器.go
# @Description :
# @Attention :
*/
package offer

// 关键是计算公式
// v=长*宽
// 然后因为只能横向移动,所以宽是必然会减少的,那么在宽减少的情况下,怎么尽量大: 通过往长的进行移动
func maxArea(height []int) int {
	r := 0
	left, right := 0, len(height)-1
	for left < right {
		// 因为水的计算,是通过短板来计算的,
		r = maxAreaMax(r, (right-left)*(maxAreaMin(height[right], height[left])))
		// 每次都往更高的板移动
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return r
}
func maxAreaMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxAreaMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
