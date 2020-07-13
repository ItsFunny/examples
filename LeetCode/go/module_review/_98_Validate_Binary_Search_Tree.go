/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-13 09:14 
# @File : _98_Validate_Binary_Search_Tree.go
# @Description : 判断是否是二叉搜索树
	什么是二叉搜索树: 左孩子 < root < 右孩子
	但是注意:
		10
	  /     \
    5         15
             /   \
	       6       20
	也是一直是 left<root<right
	所以判断是否是完全二叉树: left 的最大值 < root < right的最小值
# @Attention : 
*/
package module_review

func isValidBST(root *TreeNode) bool {
	result := validBst(root)
	return result.IsValid
}

type resultType struct {
	IsValid bool
	Max     *TreeNode
	Min     *TreeNode
}

func validBst(root *TreeNode) resultType {
	var result resultType
	if root == nil {
		result.IsValid = true
		return result
	}

	left := validBst(root.Left)
	right := validBst(root.Right)

	if !left.IsValid || !right.IsValid {
		result.IsValid = false
		return result
	}
	if (nil != left.Max && left.Max.Val >= root.Val) || (nil != right.Min && right.Min.Val <= root.Val) {
		result.IsValid = false
		return result
	}
	result.IsValid = true
	result.Min = root
	if left.Min != nil {
		result.Min = left.Min
	}
	result.Max = root
	if right.Max != nil {
		result.Max = right.Max
	}

	return result
}
