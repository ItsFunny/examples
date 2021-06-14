/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/30 2:17 下午
# @File : evaluate_reverse_polish_notation.go
# @Description :
波兰表达式
# @Attention :
*/
package v2

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens)==0{
		return 0
	}
	// 1. 遍历,然后发现如果不是运算符则入栈,是的话,则全部出栈,然后运算再入栈
	stack := make([]int, 0)
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			if len(stack)<2{
				return -1
			}
			first:=stack[len(stack)-2]
			second:=stack[len(stack)-1]
			stack=stack[:len(stack)-2]
			res:=0
			if v == "+" {
				res = first + second
			} else if v == "-" {
				res = first - second
			} else if v == "*" {
				res = first * second
			} else {
				res = first / second
			}
			stack = append(stack, res)
		default:
			intV, _ := strconv.Atoi(v)
			stack = append(stack, intV)
		}
	}
	return stack[0]
}
