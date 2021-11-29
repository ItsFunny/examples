/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/18 9:27 上午
# @File : insert.go
# @Description :
# @Attention :
*/
package sort

// 插入排序
// 关键:
// 1. 也是双重for 循环
// 2. 假设在该下标之前的全是有序的,第二重循环的目的在于,找到小于要插入值的最大值
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] >= temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
	return arr
}
