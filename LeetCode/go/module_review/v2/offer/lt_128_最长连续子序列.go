/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/11 8:59 上午
# @File : lt_128_最长连续子序列.go
# @Description :
# @Attention :
*/
package offer

// 解题关键: 用一个hashSet来处理,
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}
	ret := 0
	for k := range set {
		// 如果之前的数不存在,则可以开始统计了
		// 当前的数是必然存在的
		if !set[k-1] {
			currentN := k
			currentL := 0
			for set[currentN] {
				currentL++
				currentN++
			}
			if currentL > ret {
				ret = currentL
			}
		}
	}
	return ret
}
