/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-17 09:32 
# @File : lt_76_Minimum_Window_Substring.go
# @Description : 
# @Attention : 
*/
package slide_window

import "math"

func minWindow(s string, t string) string {
	// windows 滑动窗口
	have := make([]int, 128)
	// 目标次数
	need := make([]int, 128)
	for _, v := range t {
		need[int(v)]++
	}
	// 通过distance 变量控制
	left, right, distance, minLen, begin := 0, 0, 0, math.MaxInt64, 0
	for right < len(s) {
		v := int(s[right])
		if need[v] == 0 {
			// 说明不是目标字符
			right++
			continue
		}
		// 说明是目标字符
		if have[v] < need[v] {
			distance++
		}
		have[v]++
		right++
		for distance == len(t) {
			leftV := int(s[left])
			if right-left < minLen {
				minLen = right - left
				begin = left
			}
			// 移动左边
			// 如果左边移除的元素不在目标范围内
			if need[leftV] == 0 {
				left++
				continue
			}
			// 如果移除的元素是目标元素,并且出现次数刚好是一致
			if have[leftV] == need[leftV] {
				distance--
			}
			have[leftV]--
			left++
		}
	}
	if minLen == math.MaxInt64 {
		return ""
	}
	return s[begin : begin+minLen]
}

func minWindow2(s string, t string) string {
	need := make(map[int]int)
	have := make(map[int]int)
	for _, v := range t {
		need[int(v)]++
	}
	left, right, minLen, begin, match := 0, 0, math.MaxInt64, 0, 0

	for right < len(s) {
		rightVal := int(s[right])
		right++
		if need[rightVal]!=0 {
			have[rightVal]++
			if have[rightVal] == need[rightVal] {
				match++
			}
		}
		for match == len(need) {
			if right-left < minLen {
				minLen = right - left
				begin = left
			}
			leftVal := int(s[left])
			left++
			if need[leftVal]!=0{
				if need[leftVal] == have[leftVal] {
					match--
				}
				have[leftVal]--
			}
		}
	}
	if minLen == math.MaxInt64 {
		return ""
	}
	return s[begin : begin+minLen]
}
