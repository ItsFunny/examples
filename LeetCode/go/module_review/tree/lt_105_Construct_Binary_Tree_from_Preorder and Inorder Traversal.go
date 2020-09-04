/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-04 09:04 
# @File : lt_105_Construct_Binary_Tree_from_Preorder and Inorder Traversal.go
# @Description : 
# @Attention : 
*/
package tree

/*
	通过先序和中序创建二叉树
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0 || len(inorder)==0{
		return nil
	}
	for i, v := range inorder {
		if v == preorder[0] {
			return &TreeNode{
				Val:   v,
				Left:  buildTree(preorder[1:i+1], inorder[:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}
