/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/14 9:39 上午
# @File : lt_判断是否是有效的二叉搜索树.go
# @Description :
# @Attention :
*/
package v2

func isValidBST(root *TreeNode) bool {
	// r := make([]int, 0)
	// isValidBSTInOrder(root, &r)
	// for i := 0; i < len(r)-1; i++ {
	// 	if r[i] >= r[i+1] {
	// 		return false
	// 	}
	// }
	// return true
	divde := isValidBSTDivde(root)
	return divde.Valid
}

type Result struct {
	Valid bool
	Min   *TreeNode
	Max   *TreeNode
}

func isValidBSTDivde(root *TreeNode) Result {
	r := Result{}
	r.Valid = true
	if root == nil {
		return r
	}
	left := isValidBSTDivde(root.Left)
	right := isValidBSTDivde(root.Right)

	if !left.Valid || !right.Valid {
		r.Valid = false
		return r
	}
	if (nil != left.Max && left.Max.Val >= root.Val) || (nil != right.Min && right.Min.Val <= root.Val) {
		r.Valid = false
		return r
	}

	r.Valid = true
	r.Min = root
	if nil != left.Min {
		r.Min = left.Min
	}
	r.Max = root
	if nil != right.Max {
		r.Max = right.Max
	}

	return r
}

// 分治法

func isValidBSTInOrder(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}
	isValidBSTInOrder(root.Left, r)
	*r = append(*r, root.Val)
	isValidBSTInOrder(root.Right, r)
}
