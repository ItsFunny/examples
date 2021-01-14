/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-06 10:06 
# @File : _46_Permutations.go
# @Description :    排列组合问题
# @Attention : 
*/
package main

func Permute(nums []int) [][]int {
	result := [][]int{}
	permutateHelper(nums, 0, len(nums), func(arr []int) {
		result = append(result, arr)
	})
	return result
}

func permutateHelper(arr []int, start, end int, cb func(arr []int)) {
	if start == end {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		cb(tmp)
	} else {
		for i := start;i<end; i+=1 {
			// Keep current number fixed
			swap(arr, start, i)
			// Perform recursively
			permutateHelper(arr, start+1, end, cb)
			// Restore swap
			swap(arr, start, i)
		}
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}