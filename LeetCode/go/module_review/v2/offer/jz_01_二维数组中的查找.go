/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/28 9:33 上午
# @File : jz_01_二维数组中的查找.go
# @Description :
# @Attention :
*/
package offer

func Find(target int, array [][]int) bool {
	// write code here
	if len(array) == 0 {
		return false
	}
	for left, end := 0, len(array)-1; left < len(array[0]) && end >= 0; {
		if array[left][end] > target {
			end--
		} else if array[left][end] < target {
			left++
		} else {
			return true
		}
	}
	return false
}
