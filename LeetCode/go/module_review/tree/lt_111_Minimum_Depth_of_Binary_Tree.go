/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-07 08:20 
# @File : lt_111_Minimum_Depth_of_Binary_Tree.go
# @Description : 
# @Attention : 
*/
package tree

import "math"

/*
	树的最短深度: 递归法(分治法)
	关键在于
		1. 当left,right为nil代表的是到了叶子节点,返回1即可
		2. 需要对左孩子比较最小值,也需要对右孩子比较最小值
 */

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minD, minDepth(root.Left))
	}
	if root.Right != nil {
		minD = min(minD, minDepth(root.Right))
	}
	return minD + 1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
