/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/30 1:52 下午
# @File : min_stack.go
# @Description :
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop()—— 删除栈顶的元素。
top()—— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
# @Attention :
*/
package v2

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	r := MinStack{}
	return r
}

func (this *MinStack) Push(val int) {
	min := this.GetMin()
	if val < min {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return 1 << 31
	}
	return this.min[len(this.min)-1]
}
