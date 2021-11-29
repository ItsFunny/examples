/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/8 8:40 上午
# @File : lt_124_二叉树的最大路径和.go
# @Description :
# @Attention :
*/
package offer

// 解题关键
// 左节点的最大路径和,右节点的最大路径和, 左+根+右的最大值
// dfs
func maxPathSum(root *TreeNode) int {
	_, ret := dfsMaxPathSum(root)
	return ret
}
func dfsMaxPathSum(root *TreeNode) (int, int) {
	if root == nil {
		return 0, -(1 << 31)
	}
	left, leftAll := dfsMaxPathSum(root.Left)
	right, rightAll := dfsMaxPathSum(root.Right)

	ret := maxPathSumMax(left, right)
	ret = maxPathSumMax(0, ret+root.Val)

	retAll := maxPathSumMax(leftAll, rightAll)
	retAll = maxPathSumMax(retAll, left+right+root.Val)
	return ret, retAll
}

func maxPathSumMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
