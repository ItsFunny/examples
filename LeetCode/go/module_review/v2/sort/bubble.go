/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/18 9:12 上午
# @File : pop.go
# @Description :
# @Attention :
*/
package sort

// 冒泡排序
// 关键:
// 暴力遍历: 小的直接交换位置即可
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
