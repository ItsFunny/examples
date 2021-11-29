/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/18 1:12 下午
# @File : shell.go
# @Description :
# @Attention :
*/
package sort

// 希尔排序
// 关键: 是插入排序的优化
// 插入排序: 假设之前的都是有序的,步长为1

func shellSort(arr []int) []int {
	stride := len(arr)
	for stride != 1 {
		stride >>= 1
		for i := 0; i < stride; i += stride {
			for j := 0; j < len(arr); j += stride {
				tmp := arr[j]
				k := j - stride
				for ; k >= 0 && arr[k] > tmp; k -= stride {
					arr[k+stride] = arr[k]
				}
				arr[k+stride] = tmp
			}
		}
	}
	return arr
}
