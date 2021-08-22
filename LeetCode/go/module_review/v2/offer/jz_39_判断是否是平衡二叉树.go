/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/19 8:37 上午
# @File : jz_39_判断是否是平衡二叉树.go
# @Description :
# @Attention :
*/
package offer

// 关键: dfs
// 不能通过判断左右孩子是否为空来判断

func IsBalanced_Solution(pRoot *TreeNode) bool {
	return IsBalanced_Solution_Dfs(pRoot) != -1
}
func IsBalanced_Solution_Dfs(node *TreeNode) int {
	if nil == node {
		return 0
	}
	left := IsBalanced_Solution_Dfs(node.Left)
	right := IsBalanced_Solution_Dfs(node.Right)
	r :=left-right
	if left == -1 || right == -1 || r < 0 {
		return -1
	}
	return IsBalanced_Solution_max(left, right) + 1
}
func abs(a, b int) int {
	if a < b {
		a, b = b, a
	}
	return a - b
}
func IsBalanced_Solution_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
