/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-23 10:33 
# @File : base.go
# @Description : 
# @Attention : 
*/
package v0

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
