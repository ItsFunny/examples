/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/16 9:09 上午
# @File : jz_数字在升序数组中出现的次数.go
# @Description :
# @Attention :
*/
package offer

func GetNumberOfK(data []int, k int) int {
	// write code here
	left := 0
	right := len(data) - 1
	for left < right {
		mid := left + (right-left)>>1
		if data[mid] < k {
			left = mid + 1
		} else {
			// 注意,这里相等的时候,也要左移动,所以最终会是左边界
			right = mid
		}
	}
	leftBound := left

	left, right = 0, len(data)-1
	for left < right {
		mid := left + (right-left)>>1
		// 这里与上面不同,这里必须要右移,所以相等的时候,left 变化
		if data[mid] <= k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	rightBound := right
	return rightBound - leftBound
}
