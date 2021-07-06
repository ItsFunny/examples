/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/12 7:23 上午
# @File : lt_二叉树的最近公共祖先.go
# @Description :
# @Attention :
*/
package v2

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	leftRoot := lowestCommonAncestor(root.Left, p, q)
	rightRoot := lowestCommonAncestor(root.Right, p, q)

	if leftRoot != nil && rightRoot != nil {
		return root
	}
	if nil != leftRoot {
		return leftRoot
	}
	return rightRoot
}

