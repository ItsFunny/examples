/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 14:07 
# @File : shell_sort.go
# @Description : 
# @Attention : 
*/
package sort


func ShellSort(data []int) {
	stride := len(data)
	for stride != 1 {
		stride >>= 1
		// 分为了stride个组,则需要对着stride进行排序
		for i := 0; i < stride; i++ {
			// 进行插入排序
			for j := 0; j < len(data); j += stride {
				val := data[j]
				k := j - stride
				for ; k >= 0 && data[k] > val; k -= stride {
					data[k+stride] = data[k]
				}
				data[k+stride] = val
			}
		}
	}
}