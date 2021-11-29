/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/9 8:41 上午
# @File : lt_125_验证回文字符串.go
# @Description :
# @Attention :
*/
package offer

import "strings"

// 解题关键: 首尾双指针遍历即可
func isPalindrome(s string) bool {
	ret := ""
	for i := range s {
		if isalnum(s[i]) {
			ret += string(s[i])
		}
	}
	s = strings.ToLower(ret)
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
