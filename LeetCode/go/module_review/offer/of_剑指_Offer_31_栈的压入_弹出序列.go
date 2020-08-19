/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 16:53 
# @File : of_剑指_Offer_31_栈的压入_弹出序列.go
# @Description : 
# @Attention : 
*/
package offer

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	index := 0
	for _, v := range pushed {
		stack = append(stack, v)
		// 如果匹配则弹出进行匹配下一个
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	// 如果全弹出来了,说明全匹配了
	return len(stack) == 0
}
