/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 15:42 
# @File : of_剑指_Offer_28_对称的二叉树.go
# @Description : 判断树是否是对称的
# @Attention : 
*/
package offer

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return recurr(root.Left, root.Right)
}
func recurr(left, right *TreeNode) bool {
	if nil == left && nil == right {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return recurr(left.Left, right.Right) && recur(left.Right, right.Left)
}
