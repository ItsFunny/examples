/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-22 10:22 
# @File : base.go
# @Description : 
# @Attention : 
*/
package tree

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Data      interface{}
	LeftNode  *TreeNode
	RightNode *TreeNode
}

// 前序遍历 根左右
func preorderTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	preorderTree(node.LeftNode)
	preorderTree(node.RightNode)
}

// 非递归先序遍历
func preorderTreeWithStack(root *TreeNode) []interface{} {
	if nil == root {
		return nil
	}
	data := make([]interface{}, 0)
	nodes := make([]*TreeNode, 0)
	for nil != root && len(nodes) != 0 {
		// 先序遍历,需要先到最左孩子
		for root != nil {
			// 保存根节点,既数据信息
			data = append(data, root.Data)
			nodes = append(nodes, root)
			root = root.LeftNode
		}
		// 此时的root为nil,代表的是最左孩子的左侧,意味着到了底部,此时已经保存了左孩子A和A的根节点PA的值
		// 这个节点为根节点,并且已经没有左孩子了,nodes[len(nodes)-2] 的值为该节点的父节点
		node := nodes[len(nodes)-1]
		//  因为该节点已经数据保存完毕了,该弹出了
		nodes = nodes[:len(nodes)-1]
		// 此时已经保存了左孩子和根节点,所以需要跳到右孩子
		root = node.LeftNode
	}
	return data
}

// 中序遍历
// 左根右
func midorderTree(root *TreeNode) {
	if root == nil {
		return
	}
	// 左
	midorderTree(root.LeftNode)
	fmt.Println(root.Data)
	midorderTree(root.RightNode)
}

// 非递归中序遍历
func midorderTreeWithStack(root *TreeNode) ([]interface{}) {
	if root == nil {
		return nil
	}
	data := make([]interface{}, 0)
	stack := make([]*TreeNode, 0)
	for nil != root && len(stack) > 0 {
		// 左根右
		for nil != root {
			stack = append(stack, root.LeftNode)
			root = root.LeftNode
		}
		root = stack[len(stack)-1]
		data = append(data, root.Data)
		stack = stack[:len(stack)-1]
		root = root.RightNode
	}

	return data
}

// 后序遍历
func backorderTree(root *TreeNode) {
	if root == nil {
		return
	}
	backorderTree(root.LeftNode)
	backorderTree(root.RightNode)
	fmt.Println(root.Data)
}

// 非递归后序遍历
func backorderTreeWithStack(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	data := make([]interface{}, 0)
	stack := make([]*TreeNode, 0)

	var lastVistiNode *TreeNode
	for nil != root && len(stack) > 0 {
		for nil != root {
			stack = append(stack, root)
			root = root.LeftNode
		}

		node := stack[len(stack)-1]
		if node.RightNode == nil || node.RightNode == lastVistiNode {
			stack = stack[:len(stack)-1]
			data = append(data, node.Data)
			lastVistiNode = node
		} else {
			root = node.RightNode
		}
	}

	return data
}

func dfs(root *TreeNode) []interface{} {
	data := make([]interface{}, 0)
	stack := make([]*TreeNode, 0)

	stack = append(stack, root)
	var tempNode *TreeNode
	for len(stack) > 0 {
		tempNode = stack[len(stack)-1]
		if tempNode == nil {
			return data
		}
		data = append(data, tempNode.Data)
		if tempNode.RightNode != nil {
			stack = append(stack, tempNode.RightNode)
		}
		if tempNode.LeftNode != nil {
			stack = append(stack, tempNode.LeftNode)
		}
	}
	return data
}

func dfsWithDivide(root *TreeNode) []interface{} {

	return divide(root)
}
func divide(root *TreeNode) []interface{} {
	result := make([]interface{}, 0)
	if root == nil {
		return result
	}
	// 分
	left := divide(root.LeftNode)
	right := divide(root.RightNode)
	// 合
	result = append(result, root.Data)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func bfs(root *TreeNode) []interface{} {
	data := make([]interface{}, 0)
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)
	var temNode *TreeNode
	for len(queue) > 0 {
		temNode = queue[0]
		if temNode == nil {
			return data
		}
		// 移除队头
		queue = queue[1:]
		data = append(data, temNode.Data)
		if temNode.LeftNode != nil {
			queue = append(queue, temNode.LeftNode)
		}

		if temNode.RightNode != nil {
			queue = append(queue, temNode.RightNode)
		}
	}
	return data
}

func mergeSort(data []int) []int {
	return mSort(data)
}
func mSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	// 分
	mid := len(data) >> 1
	left := mSort(data[:mid])
	right := mSort(data[mid:])

	// 合
	result := merge(left, right)
	return result
}

// 合并就是按照元素排序
func merge(left []int, right []int) []int {
	i := 0
	j := 0
	index := 0
	data := make([]int, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			data[index] = left[i]
		} else {
			data[index] = right[j]
		}
		i++
		j++
		index++
	}
	data = append(data, left[i:]...)
	data = append(data, right[i:]...)
	return data
}

// 分治法的另一使用: 快速排序
func QuickSort(data []int) {
	quickSort(data, 0, len(data))
}
func quickSort(data []int, start, end int) {
	if start < end {
		// 分治法: divide
		divide := paration(data, start, end)
		// 对左边进行分
		quickSort(data, 0, divide-1)
		quickSort(data, divide+1, end)
	}
}
func paration(data []int, start, end int) int {
	// 以最左为基准数
	base := data[start]
	i := start
	j := end
	for i < j {
		for j > i && data[j] > base {
			j--
		}
		for i < j && data[i] < base {
			i++
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	// 交换基准数
	data[start], data[i] = data[i], base
	return i
}

// 求二叉树的最大深度 也是可以使用分治法
func treeMaxDepth(root *TreeNode) int {
	// 递归退出
	if root == nil {
		return 0
	}
	// 分
	left := treeMaxDepth(root.LeftNode)
	right := treeMaxDepth(root.RightNode)
	// 合
	if left > right {
		return left + 1
	}
	return right + 1
}

// 判断是否是一颗平衡的二叉树:
// 所有左子树平衡 && 所有右子树平衡 && 右子树-左子树<=1  && 左子树-右子树<=1
func IsBalanceTree(root *TreeNode) bool {

	if -1 == isBalanceTree(root) {
		return false
	}
	return true
}

func isBalanceTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := isBalanceTree(node.LeftNode)
	right := isBalanceTree(node.RightNode)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

// 获取一颗树的最长路径
func MaxTreeTrace(root *TreeNode) int {
	max := math.MaxInt64
	maxTreeTrace(root, &max)
	return max
}

func maxTreeTrace(node *TreeNode, i *int) int {
	// 递归退出条件
	if node == nil {
		return 0
	}
	// 分
	left := maxTreeTrace(node.LeftNode, i)
	right := maxTreeTrace(node.RightNode, i)

	// 过当前根节点的最大和
	lOrR := max(left, right)
	curSum := max(node.Data.(int), lOrR+node.Data.(int))
	// 考虑 横跨的情况,既从左子树到右子树 或者是右子树到左子树
	curMax := max(node.Data.(int), left+right+node.Data.(int))

	*i = max(curMax, *i)
	// 返回的是过当前节点的最大和
	return curSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 找到给定2个节点的公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 退出条件
	if root == nil {
		return nil
	}
	// 判断是否到了左子树
	if root == p || root == q {
		return root
	}
	// 分
	left := lowestCommonAncestor(root.LeftNode, p, q)
	right := lowestCommonAncestor(root.RightNode, p, q)

	// 合
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

func levelTree(root *TreeNode) [][]interface{} {
	if root == nil {
		return nil
	}
	result := make([][]interface{}, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		data := make([]interface{}, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			queue = queue[1:]
			data = append(data, node.Data)
			if nil != node.LeftNode {
				queue = append(queue, node.LeftNode)
			}
			if nil != node.RightNode {
				queue = append(queue, node.RightNode)
			}
		}
		result = append(result, data)
	}
	return result
}

// 从底下上的层级遍历
func levelOrderFromBottom(root *TreeNode) []interface{} {
	resutl := levelOo(root)
	reverse(&resutl)
	return resutl
}

func reverse(data *[]interface{}) {
	d := *data
	for i, j := 0, len(*data)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
}

func levelOo(node *TreeNode) []interface{} {
	if nil == node {
		return nil
	}
	result := make([]interface{}, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		if t.LeftNode != nil {
			queue = append(queue, t.LeftNode)
		}
		if t.RightNode != nil {
			queue = append(queue, t.RightNode)
		}
		result = append(result, t.Data)
	}
	return result
}
