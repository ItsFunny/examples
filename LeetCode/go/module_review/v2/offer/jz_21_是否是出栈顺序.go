/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/5 9:26 上午
# @File : lt_21_是否是出栈顺序.go
# @Description :
# @Attention :
*/
package offer

func IsPopOrder(pushV []int, popV []int) bool {
	stack := make([]int, 0)
	for _, v := range pushV {
		if v != popV[0] {
			stack = append(stack, v)
		} else {
			popV = popV[1:]
			for len(popV) > 0 && len(stack) > 0 {
				if popV[0] != stack[len(stack)-1] {
					break
				}
				if popV[0] == stack[len(stack)-1] {
					popV = popV[1:]
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	return len(stack) == 0
}
