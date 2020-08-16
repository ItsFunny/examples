/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 13:33 
# @File : select_sort.go
# @Description : 
# @Attention : 
*/
package sort

func SelectSort(data []int) {
	for i := 0; i < len(data); i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		data[minIndex], data[i] = data[i], data[minIndex]
	}
}
