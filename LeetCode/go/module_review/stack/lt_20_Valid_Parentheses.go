/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-19 10:15 
# @File : lt_20_Valid_Parentheses.go
# @Description : 
# @Attention : 
*/
package stack

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]int, 0)
	for _, v := range s {
		switch v {
		case ')', ']', '}':
			if len(stack) != 0 {
				last := stack[len(stack)-1]
				switch v {
				case ')':
					if last != '(' {
						return false
					}
				case ']':
					if last != '[' {
						return false
					}
				case '}':
					if last != '{' {
						return false
					}
				}
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, int(v))
			}
		default:
			stack = append(stack, int(v))
		}
	}
	return len(stack) == 0
}
