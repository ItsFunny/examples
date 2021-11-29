/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 8:09 上午
# @File : lt_863_二叉树中所有节点为k的值.go
# @Description :
# @Attention :
*/
package v2

// 关键: dfs
// 还有一个关键点就是 除了向下搜,还得能向上搜,就是直接通过parent搜,并且,很关键的一点是,防止走回头路
// 由于是二叉链表，所以，无法直接由当前结点走向其父节点，所以用了一个map加一次dfs遍历来保存所有结点的父节点，这样就可以查表直接跳到父节点了。
// 为了防止走回头路，所以设计了一个from标志，等效于设置一个set（将路过的结点加入set，若待访问的结点不在set中，则访问它，否则跳过）。
// 用set防止走回头路的解法，相较于设置from标志来说时间复杂度高一点，因为需要在set中进行插入和查找...
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parentsMap := make(map[int]*TreeNode)
	var fillParents func(node *TreeNode)
	fillParents = func(node *TreeNode) {
		if nil != node.Left {
			parentsMap[node.Left.Val] = node
			fillParents(node.Left)
		}
		if nil != node.Right {
			parentsMap[node.Right.Val] = node
			fillParents(node.Right)
		}
	}
	fillParents(root)
	ret := make([]int, 0)
	// 然后从target 开始dfs
	var dfsFillRet func(node *TreeNode, from *TreeNode, depth int)
	dfsFillRet = func(node *TreeNode, from *TreeNode, depth int) {
		if nil == node {
			return
		}
		if depth == k {
			ret = append(ret, node.Val)
			return
		}
		// 为什么下面的限制条件全是: 节点不能等于from呢,而不是直接判断为空即可,
		// 原因: 为了防止走回头路

		if node.Left != from {
			dfsFillRet(node.Left, node, depth+1)
		}
		if node.Right != from {
			dfsFillRet(node.Right, node, depth+1)
		}
		// 向上搜
		if pNode := parentsMap[node.Val]; pNode != from {
			dfsFillRet(pNode, node, depth+1)
		}
	}
	dfsFillRet(target, nil, 0)

	return ret
}
