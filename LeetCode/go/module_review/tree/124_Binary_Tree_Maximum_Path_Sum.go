/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-09 09:42 
# @File : 124_Binary_Tree_Maximum_Path_Sum.go
# @Description :   求树的最大路径和
	分治法
	1. 左子树的最大路径
	2. 右子树的最大路径
	3. 根节点  和 根节点 + max(left,right) 的最大值   curWithRoot
	4. curWithRoot 和 左+右+根的最大值
# @Attention :
*/
package tree

import "math"

func maxPathSum(root *TreeNode) int {
	max := root.Val

	maxPath(root, &max)
	return max
}

func maxPath(node *TreeNode, mm *int) int {
	if node == nil {
		return -1 * math.MaxInt32
	}
	left := maxPath(node.Left, mm)
	right := maxPath(node.Right, mm)

	// 计算左右子树的最大值
	leftRightMax := max(left, right)
	// 计算加上了根节点的最大值
	withRootMax := max(node.Val, leftRightMax+node.Val)
	// 计算横跨时候的最大值
	withThroughMax := max(withRootMax, left+right+node.Val)
	// 与最后的结果比较
	*mm = max(withThroughMax, *mm)

	return withRootMax
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
