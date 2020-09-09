/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-09 08:14 
# @File : lt_114_Flatten_Binary_Tree_to_Linked_List.go
# @Description : 
# @Attention : 
*/
package tree

/*
	就是将二叉树转换成排序链表
	1. 左子树,插到右子树的位置,将原先的右子树插到左子树的最右子树
	递归法解决: 主要要断开左孩子
	核心: 题目规律: 先序+递归
 */

func flatten(root *TreeNode) {
	if nil == root  {
		return
	}else if root.Left==nil{
		flatten(root.Right)
		return
	}

	preRight := root.Right
	left := root.Left
	root.Left=nil
	root.Right = left
	walkerNode := left
	for nil != walkerNode.Right {
		walkerNode = walkerNode.Right
	}
	walkerNode.Right = preRight
	flatten(root.Right)
}
