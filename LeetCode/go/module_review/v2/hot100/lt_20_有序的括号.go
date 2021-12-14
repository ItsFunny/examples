/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/14 9:37 上午
# @File : lt_20_有序的括号.go
# @Description :
# @Attention :
*/
package hot100

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 关键: 用栈处理
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			// 如果非 左边方向,则栈弹出匹配是否是相反的值
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (v == ')' && pop != '(') || (v == '}' && pop != '{') || (v == ']' && pop != '[') {
				return false
			}
		}
	}
	return len(stack) == 0
}
