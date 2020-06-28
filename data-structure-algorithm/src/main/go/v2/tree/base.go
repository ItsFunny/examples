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
