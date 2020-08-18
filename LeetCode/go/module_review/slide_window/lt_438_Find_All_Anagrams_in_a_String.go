/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-17 10:46 
# @File : lt_438_Find_All_Anagrams_in_a_String.go
# @Description : 
# @Attention : 
*/
package slide_window

func findAnagrams(s string, p string) []int {
	need := make(map[int]int)
	have := make(map[int]int)
	for _, v := range p {
		need[int(v)]++
	}
	result := make([]int, 0)
	left, right, match := 0, 0, 0
	for right < len(s) {
		rightVal := int(s[right])
		right++
		if need[rightVal] != 0 {
			have[rightVal]++
			if have[rightVal] == need[rightVal] {
				match++
			}
			for right-left >= len(p) {
				// 可能存在 大于 该字符串的可能
				if right-left == len(p) && match == len(need) {
					result = append(result, left)
				}
				leftVal := int(s[left])
				left++
				if need[leftVal] != 0 {
					if need[leftVal] == have[leftVal] {
						match--
					}
					have[leftVal]--
				}
			}
		}
	}
	return result
}
