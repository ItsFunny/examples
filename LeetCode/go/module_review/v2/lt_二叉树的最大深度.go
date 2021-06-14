/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/10 9:43 上午
# @File : lt_二叉树的最大深度.go
# @Description :
# @Attention :
*/
package v2

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
