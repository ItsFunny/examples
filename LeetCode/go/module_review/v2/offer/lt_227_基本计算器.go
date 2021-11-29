/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/17 9:29 上午
# @File : lt_227_基本计算器.go
# @Description :
# @Attention :
*/
package offer

// 思路:
// 1. 栈进行对结果进行保存,对于 * 或者 / 则直接取出进行计算
func calculate(s string) int {
	retStack := []int{}
	num := 0
	preSign := '+'
	for i, v := range s {
		isDigital := v >= '0' && v <= '9'
		if isDigital {
			num = num*10 + int(v-'0')
		}
		if !isDigital && v != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				retStack = append(retStack, num)
			case '-':
				retStack = append(retStack, -num)
			case '*':
				retStack[len(retStack)-1] *= num
			default:
				retStack[len(retStack)-1] /= num
			}
			preSign = v
			num = 0
		}
	}

	ret := 0
	for _, v := range retStack {
		ret += v
	}
	return ret
}
