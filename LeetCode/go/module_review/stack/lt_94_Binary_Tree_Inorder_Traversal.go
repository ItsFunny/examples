/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 09:47 
# @File : lt_94_Binary_Tree_Inorder_Traversal.go
# @Description : 
# @Attention : 
*/
package stack

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)

	walkerNode := root
	for nil != walkerNode || len(stack) > 0 {
		// 树的遍历用栈的时候,都需要将相关的元素入栈
		for nil != walkerNode {
			stack = append(stack, walkerNode)
			walkerNode = walkerNode.Left
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		walkerNode=val.Right
		result = append(result, val.Val)
	}

	return result
}
