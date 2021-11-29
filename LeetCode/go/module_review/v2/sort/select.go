/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/18 9:02 上午
# @File : select.go
# @Description :
# @Attention :
*/
package sort

// 选择排序
// 关键:
// 双重for 循环遍历
// 第二重for循环用途在于:把最小的那个进行交换
func SelectionSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
