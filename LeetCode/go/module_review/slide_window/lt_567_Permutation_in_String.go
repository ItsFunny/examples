/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 09:22 
# @File : lt_567_Permutation_in_String.go
# @Description : 
# @Attention : 
*/
package slide_window

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[int]int)
	have := make(map[int]int)
	for _, v := range s1 {
		need[int(v)]++
	}
	left, right, match := 0, 0, 0
	for right < len(s2) {
		rightVal := int(s2[right])
		right++
		if _, ok := need[rightVal]; ok {
			have[rightVal]++
			if have[rightVal] == need[rightVal] {
				match++
			}
			for right-left >= len(s1) {
				if match == len(need) {
					return true
				}
				leftVal := int(s2[left])
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
	return false
}
