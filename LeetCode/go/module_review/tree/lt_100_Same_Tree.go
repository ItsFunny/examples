/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-03 08:18 
# @File : lt_100_Same_Tree.go
# @Description : 
# @Attention : 
*/
package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return sameTree(p, q)
}
func sameTree(p *TreeNode, q *TreeNode) bool {
	if nil == p && nil == q {
		return true
	}
	if (nil == p && nil != q) || (nil == q && nil != p) || p.Val != q.Val {
		return false
	}
	return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
}
