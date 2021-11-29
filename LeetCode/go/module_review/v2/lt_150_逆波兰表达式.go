/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/21 8:06 上午
# @File : lt_150_逆波兰表达式.go
# @Description :
# @Attention :
*/
package v2

import "strconv"

func evalRPN150(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1
			}
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, v1+v2)
			case "-":
				stack = append(stack, v1-v2)
			case "*":
				stack = append(stack, v1*v2)
			case "/":
				stack = append(stack, v1/v2)
			}
		default:
			intV, _ := strconv.Atoi(v)
			stack = append(stack, intV)
		}
	}
	return stack[0]
}
