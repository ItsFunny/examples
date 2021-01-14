/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-07 08:10 
# @File : lt_108_Convert_Sorted_Array_to_Binary_Search_Tree.go
# @Description : 
# @Attention : 
*/
package tree

/*
	有序数组转换为二叉树
	解决关键: 将中间数字作为根节点,则左边的肯定是小于该值,而右边肯定是大于该值
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) >> 1
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
