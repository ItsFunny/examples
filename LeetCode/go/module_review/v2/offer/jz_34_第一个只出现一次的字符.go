/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/14 8:05 下午
# @File : jz_34_第一个只出现一次的字符.go
# @Description :
# @Attention :
*/
package offer

func FirstNotRepeatingChar(str string) int {
	// write code here
	bufs := make([]byte, 128)
	for i := 0; i < len(str); i++ {
		bufs[str[i]]++
	}
	for index, v := range str {
		if bufs[v] == 1 {
			return index
		}
	}
	return -1
}
