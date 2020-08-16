/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 13:16 
# @File : buble_sort.go
# @Description : 
# @Attention : 
*/
package sort

func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
