/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-21 09:06 
# @File : base.go
# @Description : 
# @Attention : 
*/
package list

type ListNode struct {
	Val  int
	Next *ListNode
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
