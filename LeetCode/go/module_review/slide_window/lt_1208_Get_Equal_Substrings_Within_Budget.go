/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-19 09:01 
# @File : lt_1208_Get_Equal_Substrings_Within_Budget.go
# @Description : 
# @Attention : 
*/
package slide_window

func equalSubstring(s string, t string, maxCost int) int {
	length:=0
	count, have, left, right := 0, maxCost, 0, 0
	for ; right < len(s); right++ {
		count += equalSubstringAbs(t[right], s[right])
		for count>have{
			count-= equalSubstringAbs(t[left], s[left])
			left++
		}
		length=equalSubstringMax(length,right-left+1)
	}
	return length
}

func equalSubstringAbs(a, b uint8) int {
	if a < b {
		return int(b - a)
	}
	return int(a - b)
}

func equalSubstringMax(a,b int)int{
	if a>b{
		return a
	}
	return b
}