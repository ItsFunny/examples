/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-06 09:10 
# @File : base.go
# @Description : 
# @Attention : 
*/
package queue

import "examples/data-structure-algorithm/src/main/go/v2/tree"

// 栈实现队列
type MyQueue struct {
	up   []int
	down []int
}

func NewMyQueue() *MyQueue {
	q := &MyQueue{
		up:   make([]int, 0),
		down: make([]int, 0),
	}
	return q
}

func (this *MyQueue) Push(data int) {
	for len(this.up) != 0 {
		last := this.up[len(this.up)-1]
		this.up = this.up[:len(this.up)-1]
		this.down = append(this.down, last)
	}
	this.up = append(this.up, data)
}

func (this *MyQueue) Pop() int {
	if len(this.down) == 0 {
		return -1
	}
	val := this.down[0]
	this.down = this.down[1:]

	return val
}

// BFS 层级遍历
func BFS(root *tree.TreeNode) []interface{} {
	if nil == root {
		return nil
	}
	result := make([]interface{}, 0)
	queue := make([]*tree.TreeNode, 0)
	queue = append(queue, root)

	for walkerNode := queue[len(queue)-1]; nil != walkerNode && len(queue) > 0; {
		val := walkerNode.Data
		queue = queue[1:]
		if walkerNode.LeftNode != nil {
			queue = append(queue, walkerNode.LeftNode)
		}
		if walkerNode.RightNode != nil {
			queue = append(queue, walkerNode.RightNode)
		}
		result = append(result, val)
	}

	return result
}
