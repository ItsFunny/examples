/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-10 08:57 
# @File : _102_Binary_Tree_Level_Order_Traversal.go
# @Description : 层次遍历二叉树,并且记录每层的相关元素
# @Attention : 
*/
package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
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
