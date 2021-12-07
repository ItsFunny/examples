/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/5 9:12 上午
# @File : offer30_包含min函数的栈.go
# @Description :
# @Attention :
*/
package offer2

import "math"

// 最好的做法是有一个单独的minStack存储了每一次的最小值
type MinStack struct {
	stack []int
	minStack []int
}

func Constructor2() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//
// type MinStack struct {
// 	stack []int
// 	min   int
// }
//
// /** initialize your data structure here. */
// func Constructor2() MinStack {
// 	return MinStack{}
// }
//
// func (this *MinStack) Push(x int) {
// 	this.stack = append(this.stack, x)
// 	if this.stack[this.min] > x {
// 		this.min = len(this.stack) - 1
// 	}
// }
//
// func (this *MinStack) Pop() {
// 	if len(this.stack) == 0 {
// 		return
// 	}
// 	l := len(this.stack) - 1
// 	this.stack = this.stack[:len(this.stack)-1]
// 	if this.min == l {
// 		this.min = 0
// 		for i := 1; i < len(this.stack)-1; i++ {
// 			if this.stack[i] < this.stack[this.min] {
// 				this.min = i
// 			}
// 		}
// 	}
// }
//
// func (this *MinStack) Top() int {
// 	if len(this.stack) == 0 {
// 		return -1
// 	}
// 	return this.stack[len(this.stack)-1]
// }
//
// func (this *MinStack) Min() int {
// 	if len(this.stack) == 0 {
// 		return -1
// 	}
// 	return this.stack[this.min]
// }
//
// /**
//  * Your MinStack object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Push(x);
//  * obj.Pop();
//  * param_3 := obj.Top();
//  * param_4 := obj.Min();
//  */
