/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-11 09:18 
# @File : lt_225_Implement_Stack_using_Queues.go
# @Description : 栈实现队列
# @Attention : 
*/
package stack

type MyQueue struct {
	Stack1 []int
	Stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	s := MyQueue{
		Stack1: make([]int, 0),
		Stack2: make([]int, 0),
	}
	return s
}

/** Push element x onto stack. */
func (this *MyQueue) Push(x int) {
	// push的时候往固定的push
	for len(this.Stack2) > 0 {
		val := this.Stack2[len(this.Stack2)-1]
		this.Stack2 = this.Stack2[:len(this.Stack2)-1]
		this.Stack1 = append(this.Stack1, val)
	}
	this.Stack1 = append(this.Stack1, x)
}

/** Removes the element on top of the stack and returns that element. */
// pop的时候也是同理,通过一个固定的pop
func (this *MyQueue) Pop() int {
	for len(this.Stack1) > 0 {
		val := this.Stack1[len(this.Stack1)-1]
		this.Stack1 = this.Stack1[:len(this.Stack1)-1]
		this.Stack2 = append(this.Stack2, val)
	}
	if len(this.Stack2) == 0 {
		return 0
	}
	val := this.Stack2[len(this.Stack2)-1]
	this.Stack2 = this.Stack2[:len(this.Stack2)-1]
	return val
}

/** Get the top element. */
func (this *MyQueue) Peek() int {
	for len(this.Stack1) > 0 {
		val := this.Stack1[len(this.Stack1)-1]
		this.Stack1 = this.Stack1[:len(this.Stack1)-1]
		this.Stack2 = append(this.Stack2, val)
	}
	if len(this.Stack2) == 0 {
		return 0
	}
	val := this.Stack2[len(this.Stack2)-1]
	return val
}

/** Returns whether the stack is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Stack2) == 0 && len(this.Stack1) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
