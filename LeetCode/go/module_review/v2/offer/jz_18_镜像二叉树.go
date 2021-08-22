/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/4 2:01 下午
# @File : lt_jz_镜像二叉树.go
# @Description :
# @Attention :
*/
package offer

// BFS ,遍历,然后每层进行交换
func Mirror(root *TreeNode) *TreeNode {
	if nil==root{
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		left := node.Left
		node.Left = node.Right
		node.Right = left
		if nil != node.Left {
			queue = append(queue, node.Left)
		}
		if nil != node.Right {
			queue = append(queue, node.Right)
		}
	}
	return root
}
