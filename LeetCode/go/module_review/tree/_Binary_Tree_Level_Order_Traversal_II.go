/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-10 09:11 
# @File : _Binary_Tree_Level_Order_Traversal_II.go
# @Description : 反转层次遍历
# @Attention : 
*/
package tree

func levelOrderBottom(root *TreeNode) [][]int {
	result := levelBfsTree(root)
	reverse(result)
	return result
}

func levelBfsTree(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

func reverse(data [][]int) {
	for i, j := 0, len(data)-1; i < j; {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
