/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 16:36 
# @File : of_剑指_Offer_30_包含min函数的栈.go
# @Description : 
# @Attention : 
*/
package offer

type MinStack struct {
	MinIndex int
	Data     []int
}

/** initialize your data structure here. */
func MMConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if x < this.Data[this.MinIndex] {
		this.MinIndex = len(this.Data) - 1
	}
}

func (this *MinStack) Pop() {
	if len(this.Data) == 0 {
		return
	}
	this.Data = this.Data[:len(this.Data)-1]
	if this.MinIndex == len(this.Data) {
		// 说明删除了最小的元素,则需要重新查询
		if len(this.Data) == 0 {
			this.MinIndex = 0
			return
		}
		min := this.Data[0]
		this.MinIndex=0
		for i := 1; i < len(this.Data); i++ {
			if this.Data[i] < min {
				this.MinIndex = i
			}
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) Min() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Data[this.MinIndex]
}
