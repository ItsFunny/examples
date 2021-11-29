/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/29 8:53 上午
# @File : lt_236_二叉树的最近公共祖先.go
# @Description :
# @Attention :
*/
package v2

// 关键: 死记硬背吧,没整理清楚
// 第二种方法是: 记录父节点的方法, 然后从p,q 往上遍历即可
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 临界条件判断
	if nil == root {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)

	if nil != left && nil != right {
		return root
	}
	if nil != right {
		return right
	}
	return left
}
