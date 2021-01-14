/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-19 10:22 
# @File : lt_150_Evaluate_Reverse_Polish_Notation2.go
# @Description : 
# @Attention : 
*/
package stack

import "strconv"

func evalRPN2(tokens []string) int {
	stack := make([]string, 0)
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			v1, _ := strconv.Atoi(stack[len(stack)-1])
			v2, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			res := 0
			switch v {
			case "+":
				res = v1 + v2
			case "-":
				res = v1 - v2
			case "*":
				res = v1 * v2
			case "/":
				res = v1 / v2
			}
			stack = append(stack, strconv.Itoa(res))
		default:
			stack = append(stack, v)
		}
	}
	v, _ := strconv.Atoi(stack[0])
	return v
}
