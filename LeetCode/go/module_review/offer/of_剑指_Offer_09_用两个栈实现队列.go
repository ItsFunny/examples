/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 11:05 
# @File : of_剑指_Offer_09_用两个栈实现队列.go
# @Description : 
# @Attention : 
*/
package offer

type CQueue struct {
	In  []int
	Out []int
}

func Constructor() CQueue {
	c := CQueue{
		In:  make([]int, 0),
		Out: make([]int, 0),
	}
	return c
}

func (this *CQueue) AppendTail(value int) {
	this.In = append(this.In, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Out) == 0 {
		if len(this.In) == 0 {
			return -1
		}
		for len(this.In) > 0 {
			v := this.In[len(this.In)-1]
			this.Out = append(this.Out, v)
			this.In = this.In[:len(this.In)-1]
		}
	}

	v := this.Out[len(this.Out)-1]
	this.Out = this.Out[:len(this.Out)-1]
	return v
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
