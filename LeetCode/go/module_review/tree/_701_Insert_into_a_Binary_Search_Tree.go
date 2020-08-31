/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-13 10:41 
# @File : _701_Insert_into_a_Binary_Search_Tree.go
# @Description : 完全二叉树值插入  前提 left<root<right
# @Attention : 
*/
package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if nil == root {
		root=&TreeNode{Val: val}
		return root
	}

	if root.Val>val {
		root.Left=insertIntoBST(root.Left,val)
	}else{
		root.Right=insertIntoBST(root.Right,val)
	}

	return root
}
