/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-07 09:20 
# @File : lt_155_Min_Stack.go
# @Description : 实现最小栈,用一个常量来代表最小值
# @Attention : 
*/
package stack

import (
	"fmt"
	"strconv"
)

type MinStack struct {
	MinIndex int
	Data     []int
	Index    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{}
	return m
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if x < this.Data[this.MinIndex] {
		this.MinIndex = this.Index
	}
	this.Index++
}

func (this *MinStack) Pop() {
	this.Index--
	fmt.Println("pop:" + strconv.Itoa(this.Data[len(this.Data)-1]))
	this.Data = this.Data[:len(this.Data)-1]
	if this.MinIndex == this.Index {
		this.MinIndex = 0
		if len(this.Data) > 0 {
			min := this.Data[0]
			for i := 1; i < len(this.Data); i++ {
				if this.Data[i] < min {
					this.MinIndex = i
					min = this.Data[i]
				}
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	return this.Data[this.MinIndex]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
