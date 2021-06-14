/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/11 9:35 上午
# @File : lt_求二叉树最大路径.go
# @Description :
# @Attention :
*/
package v2

func maxPathSum(root *TreeNode) int {
	_, i := maxPathSumDFS(root)
	return i
}
func maxPathSumDFS(node *TreeNode) (int, int) {
	if node == nil {
		return 0, -(1 << 31)
	}
	left, leftAll := maxPathSumDFS(node.Left)
	right, rightAll := maxPathSumDFS(node.Right)

	nodeMax := maxPathSumMax(left, right)
	nodeMax = maxPathSumMax(nodeMax+node.Val, 0)

	nodeAll := maxPathSumMax(leftAll, rightAll)
	nodeAll = maxPathSumMax(nodeAll, left+right+node.Val)

	return nodeMax, nodeAll
}
func maxPathSumMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
