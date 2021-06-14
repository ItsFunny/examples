/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/1 9:32 上午
# @File : binary_tree_inorder_traversal.go
# @Description :
中序遍历二叉树
# @Attention :
*/
package v2

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for nil != root {
			stack = append(stack, root)
			root = root.Left
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result=append(result,val.Val)
		root=val.Right
	}
	return result
}
