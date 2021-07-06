/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/28 9:42 上午
# @File : jz_05_两个栈实现队列.go
# @Description :
# @Attention :
*/
package offer

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) > 0 {
		v := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		return v
	}
	for len(stack1) > 0 {
		v:=stack1[len(stack1)-1]
		stack1=stack1[:len(stack1)-1]
		stack2=append(stack2,v)
	}
	v := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return v
}
