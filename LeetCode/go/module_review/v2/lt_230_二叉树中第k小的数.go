/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/28 9:37 上午
# @File : lt_230_二叉树中第k小的数.go
# @Description :
# @Attention :
*/
package v2

// 关键:
// 1. 题目关键: 限定了为二叉搜索树: 左<根<右
// 2. 中序遍历: 中序遍历二叉搜索树: 使得是从小到大排序
func kthSmallest(root *TreeNode, k int) int {
	ret := make([]int, 0)
	kthSmallestInOrderer(root, &ret)
	return ret[k-1]
}
func kthSmallestInOrderer(root *TreeNode, ret *[]int) {
	if nil != root {
		kthSmallestInOrderer(root.Left, ret)
		*ret = append(*ret, root.Val)
		kthSmallestInOrderer(root.Right, ret)
	}
}
