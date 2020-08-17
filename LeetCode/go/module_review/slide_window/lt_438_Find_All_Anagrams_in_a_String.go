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
	result:=make([]int,0)
	have := make([]int, 128)
	need := make([]int, 128)
	for _, v := range p {
		need[v]++
	}
	left, right, distance := 0, 0, 0
	for right < len(s) {
		rightVal := s[right]
		if need[rightVal] == 0 {
			right++
			continue
		}
		if have[rightVal] < need[rightVal] {
			distance++
		}
		have[rightVal]++
		rightVal++
		for distance == len(p) {
			leftVal:=s[left]
			if need[leftVal]==0{
				left++
				continue
			}
			if need[leftVal]==have[leftVal]{

			}
		}
	}

	return result
}
