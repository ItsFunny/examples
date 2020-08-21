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

// 通过分治法遍历二叉树
func loopTreeByDivide(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	left := loopTreeByDivide(root.Left)
	right := loopTreeByDivide(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)

	return result
}

// 归并排序
func mergeSort(nums []int) {

}
func divide(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	d := len(nums) >> 1
	left := divide(nums[:d])
	right := divide(nums[d:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, 0)
	i := 0
	j := 0
	for ; i < len(left) && j < len(right); {
		i++
		j++
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[i])
		i++
	}

	return result
}

// 快速排序
func QSort(data []int) {
	qSort(data, 0, len(data)-1)
}
func qSort(data []int, start int, end int) {
	if start < end {
		index := paration(data, start, end)
		qSort(data, start, index)
		qSort(data, index+1, end)
	}
}
func paration(data []int, start, end int) int {
	standard := data[start]
	for start < end {
		for end > start && data[end] >= standard {
			end--
		}
		data[start] = data[end]

		for start < end && data[start] <= standard {
			start++
		}
		data[end] = data[start]
	}
	data[start] = standard
	return start
}

func maxDepth(root *TreeNode) int {
	return dep(root)
}
func dep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dep(root.Left)
	right := dep(root.Right)
	if left < right {
		return right + 1
	}
	return left + 1
}

// 判断是否是高度平衡的二叉树
// 1. 则子树也要平衡 ,判断深度即可

func isBalanced(root *TreeNode) bool {

	return balanced(root) > -1
}

func balanced(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := balanced(node.Left)
	right := balanced(node.Right)

	if left == -1 || right == -1 || mabs(left, right) <= -1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func mabs(a, b int) int {
	if a < b {
		return a - b
	}
	return b - a
}

// 树的最大路径和
func MaxSumOfTree(root *TreeNode) int {
	return maxSumOfTree(root)
}

func maxSumOfTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxSumOfTree(node.Left)
	right := maxSumOfTree(node.Right)
	leftRightMax := max(left, right)
	sumMax := max(leftRightMax+node.Val, node.Val)
	allMax := max(sumMax, left+right+node.Val)
	return allMax
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
