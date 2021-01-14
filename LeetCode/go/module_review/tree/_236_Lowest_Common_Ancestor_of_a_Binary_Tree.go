/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-09 10:09 
# @File : _236_Lowest_Common_Ancestor_of_a_Binary_Tree.go
# @Description : 找到2个节点的最近公共祖先
	分治法
	1.
# @Attention : 
*/
package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 从根节点开始遍历
	if nil == root {
		return nil
	}
	// 判断是否到了当前需要判断的节点
	if root == p || root == q {
		return root
	}
	// 分
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左右都不为空,说明刚好这个节点是直属root节点
	if left != nil && right != nil {
		return root
	}

	if nil != left {
		return left
	}
	if nil != right {
		return right
	}
	return nil
}
