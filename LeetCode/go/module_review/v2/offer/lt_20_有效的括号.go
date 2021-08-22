/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/21 9:15 下午
# @File : lt_20_有效的括号.go
# @Description :
# @Attention :
*/
package offer

func isValid(s string) bool {
	if len(s)&1 != 0 {
		return false
	}
	m := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	stack := make([]byte, 0)
	for index := range s {
		v := s[index]
		switch v {
		case '{', '(', '[':
			stack = append(stack, v)
		default:
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pop != m[v] {
				return false
			}
		}
	}
	return len(stack) == 0
}
