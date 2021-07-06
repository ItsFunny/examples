/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/6 9:06 上午
# @File : jz_22_从上到下打印树.go
# @Description :
# @Attention :
*/
package offer

func PrintFromTopToBottom(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	r := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		r = append(r, node.Val)
		if nil != node.Left {
			queue = append(queue, node.Left)
		}
		if nil != node.Right {
			queue = append(queue, node.Right)
		}
	}
	return r
}
