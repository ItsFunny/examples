/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-05 12:17 
# @File : _3_Longest_Substring_Without_Repeating_Characters.go
# @Description :
	Given a string,
	find the length of the longest substring without repeating characters.
	最长子串
	既双指针
# @Attention : 
*/
package main

import (
	"fmt"
)
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func lengthOfLongestSubstring(s string) int {
	maxLen, start := 0, 0
	table := [128]int{}
	for i, _ := range table {
		table[i] = -1
	}
	for i, c := range s {
		if table[c] >= start {
			start = table[c] + 1
		}
		table[c] = i
		maxLen = maxInt(maxLen, i - start + 1)
	}
	return maxLen
}
func main() {
	fmt.Println(lengthOfLongestSubstring("asd"))
}
