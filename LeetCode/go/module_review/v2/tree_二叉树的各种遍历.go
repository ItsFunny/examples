/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/6 11:19 上午
# @File : tree.go
# @Description :
# @Attention :
*/
package v2

import "fmt"

// 递归先序遍历
// 根左右
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 递归中序遍历
// 左根右
func inorderLoopTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderLoopTraversal(root.Left)
	fmt.Println(root.Val)
	inorderLoopTraversal(root.Right)
}

// 递归后序遍历
func afterOrderLoopTraversal(root *TreeNode) {
	if nil == root {
		return
	}
	afterOrderLoopTraversal(root.Left)
	afterOrderLoopTraversal(root.Right)
	fmt.Println(root.Val)
}

// 非递归先序遍历
func preOrderStackTree(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	r := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for nil != root {
			r = append(r, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return r
}

// 非递归中序遍历
func inorderStackTree(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	stack := make([]*TreeNode, 0)
	r := make([]int, 0)
	for nil != root || len(stack) > 0 {
		for nil != root {
			// 因为中序遍历是 左根右,所以要先把所有的左节点入栈
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r = append(r, node.Val)
		root = node.Right
	}

	return r
}

// 后序非递归遍历
func afterStackTree(root *TreeNode) []int {
	if nil != root {
		return nil
	}

	stack := make([]*TreeNode, 0)
	r := make([]int, 0)

	var lastVisit *TreeNode
	for nil != root || len(stack) > 0 {
		for nil != root {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		// 中序遍历导致根节点必须在右节点之后,所以必须要等待这个节点的右节点已经访问过了才可以继续
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			r = append(r, node.Val)
			lastVisit = node
		} else {
			node = node.Right
		}
		stack = stack[:len(stack)-1]
		r = append(r, node.Val)
	}

	return r
}

// bfs层次遍历
// 层次遍历都是用queue
func bfsTree(root *TreeNode) [][]int {
	r := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			list = append(list, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		r = append(r, list)
	}

	return r
}

// 递归版dfs,从上到下
func dfsLoopTopDown(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfsLoopTopDown(root.Left, result)
	dfsLoopTopDown(root.Right, result)
}

// 递归版dfs,从下到上
func dfsLoopDownTop(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := dfsLoopDownTop(root.Left)
	right := dfsLoopDownTop(root.Right)
	r := make([]int, 0)
	r = append(r, root.Val)
	r = append(r, left...)
	r = append(r, right...)
	return r
}

// 非递归版dfs
// 关键: 右节点先入栈
func dfsStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		r = append(r, node.Val)
	}
	return r
}
