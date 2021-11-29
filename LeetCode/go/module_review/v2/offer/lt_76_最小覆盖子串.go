/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/7 9:22 下午
# @File : lt_76_最小覆盖子串.go
# @Description :
# @Attention :
*/
package offer

import "math"

func minWindow(s string, t string) string {
	left, right := 0, 0

	match := 0
	start := 0
	end := 0
	min := math.MaxInt32
	need := make(map[byte]int)
	have := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	for left<right {
		c := s[right]
		right++
		if need[c] != 0 {
			have[c]++
			if have[c] == need[c] {
				match++
			}
		}
		for match == len(t) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c := s[left]
			// 左指针右移
			left++
			if need[c] != 0 {
				if have[c] == need[c] {
					match--
				}
				have[c]--
			}
		}
	}
	if min == math.MaxInt32 {
		return ""
	}

	return s[start:end]
}
