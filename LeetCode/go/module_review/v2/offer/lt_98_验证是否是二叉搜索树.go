/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/17 8:51 上午
# @File : lt_98_验证是否是二叉搜索树.go
# @Description :
# @Attention :
*/
package offer

// 解题关键
// 根节点> 左节点 && 根节点< 右节点
// 左边的最大值要小于root 要小于右边的最小值
func isValidBST(root *TreeNode) bool {
	return dfsIsValidBST(root).valid
}

type TempResult struct {
	valid bool
	max   *TreeNode
	min   *TreeNode
}

func dfsIsValidBST(node *TreeNode) TempResult {
	r := TempResult{}
	r.valid = true
	if nil == node {
		return r
	}
	left := dfsIsValidBST(node.Left)
	right := dfsIsValidBST(node.Right)

	if !left.valid || !right.valid {
		r.valid = false
		return r
	}
	// 左边的最大值要小于root<右边最小值
	if (nil != left.max && left.max.Val >= node.Val) || (nil != right.min && right.min.Val <= node.Val) {
		r.valid = false
		return r
	}
	r.max = node
	if nil != right.max {
		r.max = right.max
	}
	r.min = node
	if nil != left.min {
		r.min = left.min
	}
	return r
}
