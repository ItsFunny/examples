/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-03 08:23 
# @File : lt_101_Symmetric_Tree.go
# @Description : 判断是否是镜像树
# @Attention : 
*/
package v0

func isSymmetric(root *TreeNode) bool {
	if nil==root{
		return true
	}
	return symmetirc(root.Left, root.Right)
}
func symmetirc(p *TreeNode, q *TreeNode) bool {
	if nil == p && nil == q {
		return true
	}
	if (nil == p && nil != q) || (nil != p && nil == q) || p.Val != q.Val {
		return false
	}
	return symmetirc(p.Left, q.Right) && symmetirc(p.Right,q.Left)
}
