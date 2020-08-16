/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 15:44 
# @File : quick_sort.go
# @Description : 
# @Attention : 
*/
package sort

func QuickSort(data []int) {
	qSort(data, 0, len(data)-1)
}
func qSort(data []int, start, end int) {
	if start < end {
		paration := paration(data, start, end)
		qSort(data, start, paration)
		qSort(data, paration+1, end)
	}
}

func paration(data []int, start, end int) int {
	standard := data[start]
	for start < end {
		for end > start && data[end] >= standard {
			end--
		}
		// 右边的值都是大于左边的,因此一旦有值小于标准则,则需要将其换到左边去,同时这个时候左边的值是刚好
		// 是临界值,既下一个可能就大于这个标准值了
		data[start] = data[end]
		for start < end && data[start] <= standard {
			start++
		}
		// 同理,左边的值都是小于右边的,一旦不满足,需要将右边的值用左边的代替
		data[end] = data[start]
	}
	data[start] = standard
	return start
}
