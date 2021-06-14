/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/31 7:19 下午
# @File : decode_string.go
# @Description :
# @Attention :
*/
package v2

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for index, _ := range s {
		if s[index] != ']' {
			stack = append(stack, s[index])
		} else {
			// 说明需要弹出元素,并且弹出的元素直到 为 [ 为止
			// 2[1,2,3]  =>
			tempStack := make([]byte, 0)
			for stack[len(stack)-1] != '[' {
				tempStack = append(tempStack, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			// 把'[' 弹出去
			stack = stack[:len(stack)-1]
			value := ""
			for i := len(tempStack) - 1; i >= 0; i-- {
				value += string(tempStack[i])
			}
			// 可能 数字是多个 数字的如100 ,或者是1000 ,但是数字的上一个必然是 ]
			// 还是需要反转
			tempStack = tempStack[:0]
			for len(stack) > 0 && stack[len(stack)-1] != ']' && stack[len(stack)-1] <= '9' {
				tempStack = append(tempStack, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			counts := reverseStack(tempStack)
			times, _ := strconv.Atoi(counts)
			value = strings.Repeat(value, times)
			for i := 0; i < len(value); i++ {
				stack = append(stack, value[i])
			}
		}
	}
	return string(stack)
}

func reverseStack(s []byte) string {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	r := ""
	for _, v := range s {
		r += string(v)
	}
	return r
}
