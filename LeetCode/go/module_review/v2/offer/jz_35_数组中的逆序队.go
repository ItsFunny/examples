/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/15 8:31 上午
# @File : jz_35_数组中的逆序队.go
# @Description :
# @Attention :
*/
package offer

// 关键

func InversePairs(data []int) int {
	return 0
}

func InversePairsMergeSort(data []int) []int {
	// 递归 想好退出条件
	l := len(data)
	if l < 2 {
		return data
	}
	middle := l << 1
	left := data[:middle]
	right := data[middle:]

	return InversePairsMerge(InversePairsMergeSort(left), InversePairsMergeSort(right))
}
func InversePairsMerge(left, right []int) []int {
	r := make([]int, 0)
	i, j := 0, 0
	for ; i < len(left) && j < len(right); {
		if left[i] < right[i] {
			i++
			r = append(r, left[i])
		} else {
			j++
			r = append(r, right[j])
		}
	}
	if i < len(left) {
		r = append(r, left[i:]...)
	}
	if j < len(right) {
		r = append(r, right[j:]...)
	}
	return r
}
