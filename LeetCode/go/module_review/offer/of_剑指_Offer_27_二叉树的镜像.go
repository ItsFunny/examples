/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 15:38 
# @File : of_剑指_Offer_27_二叉树的镜像.go
# @Description : 
# @Attention : 
*/
package offer

func mirrorTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if nil != node.Left {
			stack = append(stack, node.Left)
		}
		if nil != node.Right {
			stack = append(stack, node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}
