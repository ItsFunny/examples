/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/7 9:47 上午
# @File : jz_24_二叉树中某个值的路径.go
# @Description :
# @Attention :
*/
package offer

func FindPath(root *TreeNode, expectNumber int) [][]int {
	r := make([][]int, 0)
	dfsFindPath(root, expectNumber, &r, make([]int, 0))
	return r
}

func dfsFindPath(root *TreeNode, left int, res *[][]int, path []int) {
	if nil == root {
		return
	}
	val := root.Val
	path = append(path, val)
	left -= val
	if left < 0 {
		return
	}
	if left == 0 && root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	dfsFindPath(root.Left, left, res, path)
	dfsFindPath(root.Right, left, res, path)
}
