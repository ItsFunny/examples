/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/17 9:34 下午
# @File : lt_162_寻找峰值.go
# @Description :
# @Attention :
*/
package offer

import "math"

// 关键:  想象为爬坡 ,
// 1. 题目关键: 峰值为 大于相邻左边和右边的值
// 二分法: 如果中间的元素 M  :  l<m<r ,表明应该往右移动
// 如果中间的元素 : l>m>r,则表明应该往左移动
func findPeakElement(nums []int) int {
	n := len(nums)
	getV := func(index int) int {
		if index == -1 || index == n {
			return math.MinInt64
		}
		return nums[index]
	}
	left, right := 0, n-1
	for {
		mid := (left + right) >> 1
		if getV(mid) > getV(mid-1) && getV(mid) > getV(mid+1) {
			return mid
		}
		if getV(mid+1) > getV(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
