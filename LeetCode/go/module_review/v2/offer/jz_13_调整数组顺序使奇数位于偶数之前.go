/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/1 9:02 上午
# @File : jz_13_调整数组顺序使奇数位于偶数之前.go
# @Description :
# @Attention :
*/
package offer

func reOrderArray(array []int) []int {
	odd := make([]int, 0)
	even := make([]int, 0)
	for _, v := range array {
		if v&1 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	odd = append(odd, even...)
	return odd
}

// 或者双指针

func reOrderArray2(array []int) []int {
	r := make([]int, len(array))
	indexHead := 0
	indexTail := len(array) - 1
	for head, last := 0, len(array)-1; head < len(array) && last >= 0; {
		if array[head]&1 > 0 {
			r[indexHead] = array[head]
			indexHead++
		}
		head++
		if array[last]&1 == 0 {
			r[indexTail] = array[last]
			indexTail--
		}
		last--
	}
	return r
}
