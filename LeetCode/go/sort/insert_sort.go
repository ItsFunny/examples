/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 13:40 
# @File : insert_sort.go
# @Description : 插入排序
# @Attention : 
*/
package sort

func InsertSort(data []int) {
	for i := 1; i < len(data); i++ {
		val := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > val; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = val
	}
}
