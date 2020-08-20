/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-20 09:44 
# @File : base.go
# @Description : 
# @Attention : 
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTree(node *TreeNode) []int {
	// 根左右
	if node == nil {
		return nil
	}
	result := make([]int, 0)
	result = append(result, node.Val)
	left := preOrderTree(node.Left)
	right := preOrderTree(node.Right)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// 前序非递归

func preOrderWithStack(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for nil != root || len(stack) > 0 {
		for nil != root {
			result = append(result, root.Val)
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序遍历
func inOrderTree(root *TreeNode) []int {
	// 左根右
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	left := inOrderTree(root.Left)
	result = append(result, root.Val)
	right := inOrderTree(root.Right)
	result = append(result, left...)
	result = append(result, right...)

	return result
}

func inOrderWithStack(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for nil != root || len(stack) > 0 {
		for nil != root {
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}

	return result
}

// 后序遍历
func afterOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	left := afterOrder(root.Left)
	right := afterOrder(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)

	return result
}

// 后序遍历不同在于 根节点必须在右节点之后弹出
func afterOrderWithStack(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	var lastNode *TreeNode
	for nil != root || len(stack) > 0 {
		for nil != root {
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		// 校验 右节点是否被弹出过了
		if node.Right == nil || node.Right == lastNode {
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			lastNode = node
		} else {
			root = node.Right
		}
	}

	return result
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([]int, 0)
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		result = append(result, node.Val)
		if nil != node.Right {
			queue = append(queue, node.Right)
		}
		if nil != node.Left {
			queue = append(queue, node.Left)
		}
		queue = queue [1:]
	}

	return result
}

func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	result := make([]int, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if nil != node.Left {
			stack = append(stack, node.Left)
		}
		if nil != node.Right {
			stack = append(stack, node.Right)
		}
		return result
	}
	return result
}
