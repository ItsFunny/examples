/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-23 09:59 
# @File : _104_Maximum_Depth_of_Binary_Tree.go
# @Description :
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:    求一颗二叉树的深度
# @Attention :  分治法
*/
package module_review

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
