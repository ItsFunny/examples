/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 16:41 
# @File : lt_344_Reverse_String.go
# @Description : 
# @Attention : 
*/
package recusive

func reverseString(s []byte) {
	result := make([]byte, 0)
	revrse(s, &result, len(s)-1)
	for i := 0; i < len(s); i++ {
		s[i] = result[i]
	}
}

func revrse(s []byte, result *[]byte, index int) {
	if index < 0 {
		return
	}
	*result = append(*result, s[index])
	revrse(s, result, index-1)
}
