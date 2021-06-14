/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-09 09:31 
# @File : _110_Balanced_Binary_Tree.go
# @Description :    判断是否是一颗平衡二叉树
	1. 左右子树高度差不超过1
// 左边需要平衡 && 右边需要平衡
# @Attention :
注意 ,当为nil的时候, 是平衡的,意味着是true
*/
package v0

func isBalanced(root *TreeNode) bool {
	b, _ := balanced(root)
	return b
}

func balanced(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	left, lh := balanced(root.Left)
	right, rh := balanced(root.Right)
	if !left || !right || (lh-rh) > 1 || (rh-lh) > 1 {
		return false, 0
	}
	if lh > rh {
		return true, lh + 1
	}
	return true, rh + 1
}
