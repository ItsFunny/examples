/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/2 9:11 上午
# @File : lt_3_最长不重复子串.go
# @Description :
# @Attention :
*/
package offer

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	r := 1
	// 左右指针
	left, right := 0, 0
	set := make(map[byte]int)
	for ; right < len(s); right++ {
		if index, exist := set[s[right]]; exist {
			// 说明有重复的了,则开始移动左指针
			left = lengthOfLongestSubstringMax(left, index)
		}
		r = lengthOfLongestSubstringMax(r, right-left+1)
		// 更新最大值
		set[s[right]] = right + 1
	}
	return r
}
func lengthOfLongestSubstringMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
