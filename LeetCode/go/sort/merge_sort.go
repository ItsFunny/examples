/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 14:34 
# @File : merge_sort.go
# @Description : 归并排序
# @Attention : 
*/
package sort

func MergeSort(data []int) []int{
	return mergeSort(data)
}

func mergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	mid := len(data) >> 1
	left := mergeSort(data[:mid])
	right := mergeSort(data[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i, j := 0, 0
	result := make([]int, 0)
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}
