/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/14 10:13 上午
# @File : lt_二叉树中插入值.go
# @Description :
//  二叉搜索树中插入
# @Attention :
*/
package v2

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
