/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/18 11:31 上午
# @File : quick.go
# @Description :
# @Attention :
*/
package sort

// 快排
// 关键
// 1. 找分割点: paration: 通过设定 standard (通常是左边这个节点作为paration,然后从右边开始找到小于标准值的),然后替换
// 2. 递归继续qsort
func quickSort(arr []int) []int {
	qSort(arr, 0, len(arr)-1)
	return arr
}
func qSort(arr []int, start, end int) {
	if start < end {
		paration := qSortParation(arr, start, end)
		qSort(arr, start, paration)
		qSort(arr, paration+1, end)
	}
}
func qSortParation(arr []int, start, end int) int {
	standard := arr[start]
	for start != end {
		for arr[end] > standard && end >= start {
			end--
		}
		arr[start] = arr[end]
		for start < end && arr[start] <= standard {
			start++
		}
		arr[end] = arr[start]
	}
	arr[start] = standard
	return start
}
