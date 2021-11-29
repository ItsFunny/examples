/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/23 8:56 上午
# @File : lt_105_先序和中序构建二叉树.go
# @Description :
# @Attention :
*/
package offer

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0 || len(inorder)==0{
		return nil
	}
	for index := range inorder {
		if preorder[0] == inorder[index] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:len(preorder)+1], inorder[0:index]),
				Right: buildTree(preorder[len(preorder)+1:], inorder[index+1:]),
			}
		}
	}
	return nil
}
