/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-29 08:38 
# @File : stack.go
# @Description : 
# @Attention : 
*/
package stack_queue

import (
	"github.com/chanxuehong/util/math"
	"strconv"
)

// 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈
// 思路: 将第0位的数设置为最小的值
type MinStack struct {
	Data     []int
	MinIndex int
	Index    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{
		Data:     make([]int, 8),
		MinIndex: 0,
	}
	m.Data[m.MinIndex] = math.MaxInt
	return m
}

func (this *MinStack) Push(x int) {
	this.ensureCap()
	this.Data[this.Index] = x
	if x < this.Data[this.MinIndex] {
		this.MinIndex = this.Index
	}
	this.Index++
}

func (this *MinStack) Pop() {
	this.Data = this.Data[:len(this.Data)-1]
	this.Index--
	if this.Index != this.MinIndex {
		return
	}
	min := math.MaxInt
	for i := 0; i < len(this.Data); i++ {
		if this.Data[i] < min {
			this.MinIndex = i
			min = this.Data[i]
		}
	}
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	return this.Data[this.MinIndex]
}

func (this *MinStack) ensureCap() {
	if this.Index >= len(this.Data)>>1 {
		this.grow()
	}
}

func (this *MinStack) grow() {
	oldCap := cap(this.Data)
	newCap := oldCap << 1
	newData := make([]int, newCap)
	copy(this.Data, newData)
	this.Data = newData
}

// //////////////////////////////////////////////////////////////////////////////////

// 波兰表达式计算 > 输入: ["2", "1", "+", "3", "*"] > 输出: 9 解释: ((2 + 1) * 3) = 9
// 思路: 栈保存元素,并且当数量为2的时候代表可以运算,计算结果之后重新推入

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {

		switch tokens[i] {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0
			}
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var result int
			switch tokens[i] {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				result = left / right
			default:
				v, _ := strconv.Atoi(tokens[i])
				stack = append(stack, v)
			}
			stack = append(stack, result)
		}
	}

	return stack[0]
}
