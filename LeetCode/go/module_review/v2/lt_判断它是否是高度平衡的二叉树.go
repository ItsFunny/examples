/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/10 10:07 上午
# @File : lt_判断它是否是高度平衡的二叉树.go
# @Description :
# @Attention :
*/
package v2

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalancedMaxDepth(root) != -1
}
func isBalancedMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := isBalancedMaxDepth(root.Left)
	right := isBalancedMaxDepth(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
