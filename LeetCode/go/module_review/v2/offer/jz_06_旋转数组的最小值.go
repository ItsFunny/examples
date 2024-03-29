/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/29 8:51 上午
# @File : jz_06_旋转数组的最小值.go
# @Description :
# @Attention :
*/
package offer

func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	if len(rotateArray) == 0 {
		return 0
	}
	start := 0
	end := len(rotateArray) - 1
	for start+1 < end {
		mid := start + (end-start)>>1
		if rotateArray[mid] > rotateArray[end] {
			// 说明翻转的地方在右边,所以需要start移动
			start = mid
		} else if rotateArray[mid] <= rotateArray[end] {
			// 说明当前mid已经在翻转里了,所以往左移动即可
			end = mid
		}
	}
	if rotateArray[start] < rotateArray[end] {
		return rotateArray[start]
	}
	return rotateArray[end]
}
