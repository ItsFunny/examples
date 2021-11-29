/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/18 10:43 上午
# @File : merge.go
# @Description :
# @Attention :
*/
package sort

// 关键:
// 1. 分治法
//

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) >> 1
	leftA := mergeSort(arr[:mid])
	rightA := mergeSort(arr[mid:])
	return merge(leftA, rightA)
}
func merge(l1, l2 []int) []int {
	ret := make([]int, 0)
	i1, i2 := 0, 0
	for ; i1 < len(l1) && i2 < len(l2); {
		if l1[i1] < l2[i2] {
			ret = append(ret, l1[i1])
			i1++
		} else {
			ret = append(ret, l2[i2])
			i2++
		}
	}
	for ; i1 < len(l1); i1++ {
		ret = append(ret, l1[i1])
	}
	for ; i2 < len(l2); i2++ {
		ret = append(ret, l2[i2])
	}
	return ret
}
