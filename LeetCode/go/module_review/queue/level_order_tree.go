/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-11 09:46 
# @File : level_order_tree.go
# @Description : 
# @Attention : 
*/
package queue

import (
	"examples/LeetCode/go/module_review/base"
)

// 层次遍历tree 既 BFS

func bfs(root *base.TreeNode) [][]int {
	if nil == root {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*base.TreeNode, 0)
	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Data)
			if nil != node.LeftNode {
				queue = append(queue, node.LeftNode)
			}
			if nil != node.RightNode {
				queue = append(queue, node.RightNode)
			}
		}
		result = append(result, list)
	}
	return result
}
