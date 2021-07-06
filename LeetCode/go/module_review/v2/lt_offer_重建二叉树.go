/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/26 4:17 下午
# @File : lt_offer_重建二叉树.go
# @Description :
# @Attention :
*/
package v2

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0 || len(inorder)==0{
		return nil
	}
	for index,v := range inorder {
		if preorder[0] == v {
			return &TreeNode{
				Val:   v,
				// 当匹配到根节点之后,左边的是左子树,右边的是右子树
				Left:  buildTree(preorder[1:index+1], inorder[0:index]),
				Right: buildTree(preorder[index+1:], inorder[index+1:]),
			}
		}
	}
	return nil
}
