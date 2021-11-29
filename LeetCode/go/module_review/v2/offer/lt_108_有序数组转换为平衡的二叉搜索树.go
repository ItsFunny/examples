/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/24 8:57 上午
# @File : lt_108_有序数组转换为平衡的二叉搜索树.go
# @Description :
# @Attention :
*/
package offer

// 解题关键
// 中间节点作为根节点
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}
func sortedArrayToBSTHelper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (right + left) >> 1
	ret := &TreeNode{Val: nums[mid]}
	ret.Left = sortedArrayToBSTHelper(nums, left, mid-1)
	ret.Right = sortedArrayToBSTHelper(nums, mid+1, right)
	return ret
}
