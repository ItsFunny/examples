/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/3 10:31 下午
# @File : implement_queue_using_stacks.go
# @Description :
//  栈实现队列
# @Attention :
*/
package v2

type MyQueue struct {
	pushStack []int
	popStack  []int
}

/** Initialize your data structure here. */
func QueueConstructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.popStack) == 0 {
		if len(this.pushStack) == 0 {
			return 0
		}
		for len(this.pushStack) > 0 {
			p := this.pushStack[len(this.pushStack)-1]
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
			this.popStack = append(this.popStack, p)
		}
	}
	v := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		if len(this.pushStack) == 0 {
			return 0
		}
		for len(this.pushStack) > 0 {
			p := this.pushStack[len(this.pushStack)-1]
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
			this.popStack = append(this.popStack, p)
		}
	}
	return this.popStack[len(this.popStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
