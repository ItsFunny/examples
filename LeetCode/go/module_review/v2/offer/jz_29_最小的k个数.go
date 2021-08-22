/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/12 9:47 上午
# @File : jz_29_最小的k个数.go
# @Description :
# @Attention :
*/
package offer

func GetLeastNumbers_Solution(input []int, k int) []int {
	heap := make([]int, 0)
	for _, v := range input {
		if len(heap) < k {
			heap = append(heap, v)
		} else {
			if v < heap[0] {
				heap = heap[1:]
				heap = append(heap, v)
			}
		}
	}
	return heap
}
