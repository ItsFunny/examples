/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/15 8:44 上午
# @File : lt_94_二叉树的中序遍历.go
# @Description :
# @Attention :
*/
package offer

// 非递归法: 既层次遍历: BFS:使用栈
func inorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for nil != root {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		root = node.Right
	}
	return ret
}

// 中序遍历方式:  左根右
func inorderTraversalWithLoop(root *TreeNode) []int {
	ret := make([]int, 0)
	if nil == root {
		return nil
	}
	loopInorderTraversal(root, &ret)
	return ret
}

// 递归法
func loopInorderTraversal(node *TreeNode, ret *[]int) {
	if nil != node.Left {
		loopInorderTraversal(node.Left, ret)
	}
	*ret = append(*ret, node.Val)
	if nil != node.Right {
		loopInorderTraversal(node.Right, ret)
	}
}
