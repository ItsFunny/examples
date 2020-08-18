/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 10:31 
# @File : of_剑指_Offer_07_重建二叉树.go
# @Description :
	输入某二叉树的前序遍历和中序遍历的结果，
请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
# @Attention :
	就是类似于不停的构建子树
*/
package offer

func buildTree2(preorder []int, inorder []int) *TreeNode {

	for index := range inorder {
		if preorder[0] == inorder[index] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree2(preorder[1:index+1], inorder[0:index]),
				Right: buildTree2(preorder[index+1:], inorder[index+1:]),
			}
		}
	}
	return nil
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	inorderIndex := 0
	l := len(preorder)
	for i := 0; i < l; i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{
				Val: preorder[i],
			}
			stack = append(stack, node.Left)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[len(stack)-1:]
				inorderIndex++
			}
		}
		node.Right = &TreeNode{
			Val: preorderVal,
		}
		stack = append(stack, node.Left)
	}

	return root
}
