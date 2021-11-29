/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/16 9:01 上午
# @File : lt_103_二叉树的锯齿形层次遍历.go
# @Description :
# @Attention :
*/
package offer

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		vals := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if l%2 == 0 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ret = append(ret, vals)
	}
	return ret
}
