/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/5 9:08 上午
# @File : offer09_栈实现队列.go
# @Description :
# @Attention :
*/
package offer2

// 关键: 一个栈push,一个栈pop
type CQueue struct {
	push []int
	pop  []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.push = append(this.push, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.pop) == 0 {
		for len(this.push) > 0 {
			v := this.push[len(this.push)-1]
			this.push = this.push[:len(this.push)-1]
			this.pop = append(this.pop, v)
		}
	}
	if len(this.pop) == 0 {
		return -1
	}
	ret := this.pop[len(this.pop)-1]
	this.pop = this.pop[:len(this.pop)-1]
	return ret
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
