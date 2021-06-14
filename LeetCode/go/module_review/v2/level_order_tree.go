/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/4 8:56 上午
# @File : level_order_tree.go
# @Description : 二叉树层次遍历,既BFS
# @Attention :
*/
package v2

func levelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	res := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			parent := queue[0]
			queue = queue[1:]
			list = append(list, parent.Val)
			if parent.Left != nil {
				queue = append(queue, parent.Left)
			}
			if parent.Right != nil {
				queue = append(queue, parent.Right)
			}
		}
		res=append(res,list)
	}

	return res
}
