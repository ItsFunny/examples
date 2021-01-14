/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 09:57 
# @File : lt_3_Longest_Substring_Without_Repeating_Characters.go
# @Description : 不重复最长子串
# @Attention : 
*/
package slide_window

func lengthOfLongestSubstring(s string) int {
	have := make(map[byte]int)

	left := 0
	right := 0
	maxLen := 0

	for right < len(s) {
		rightVal := s[right]
		right++
		have[rightVal]++
		// 当有重复的数字之后
		for have[rightVal] > 1 {
			leftVal := s[left]
			left++
			have[leftVal]--
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}

	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	// 滑动窗口核心点：1、右指针右移 2、根据题意收缩窗口 3、左指针右移更新窗口 4、根据题意计算结果
	if len(s) == 0 {
		return 0
	}
	win := make(map[byte]int)
	left := 0
	right := 0
	ans := 1
	for right < len(s) {
		c := s[right]
		right++
		win[c]++
		// 缩小窗口
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		// 计算结果
		ans = max(right-left, ans)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
