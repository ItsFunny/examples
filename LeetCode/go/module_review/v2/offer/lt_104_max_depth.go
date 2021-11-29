/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/22 9:08 上午
# @File : lt_104_max_depth.go
# @Description :
# @Attention :
*/
package offer

func maxDepth(root *TreeNode) int {
	if  root==nil{
		return 0
	}
	left:=maxDepth(root.Left)
	right:=maxDepth(root.Right)
	if left>right{
		return left+1
	}
	return right+1
}