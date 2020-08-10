/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-07 09:51 
# @File : lt_150_Evaluate_Reverse_Polish_Notation.go
# @Description : 
# @Attention : 
*/
package stack

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	if nil == tokens || len(tokens) == 0 {
		return 0
	}
	stack := make([]string, 0)
	var result int

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			value0, _ := strconv.Atoi(stack[len(stack)-2])
			value1, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "+":
				result = value0 + value1
			case "-":
				result = value0 - value1
			case "*":
				result = value0 * value1
			case "/":
				result = value0 / value1
			}
			stack = append(stack, strconv.Itoa(result))
		default:
			stack = append(stack, tokens[i])
		}

	}

	res, _ := strconv.Atoi(stack[0])
	return res
}
