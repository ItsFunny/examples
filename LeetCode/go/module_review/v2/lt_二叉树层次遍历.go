/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/12 7:49 上午
# @File : lt_二叉树层次遍历.go
# @Description :
# @Attention :
*/
package v2

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	r := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue=queue[1:]
			list = append(list, node.Val)
			if nil != node.Left {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
		}
		r=append(r,list)
	}
	return r
}
