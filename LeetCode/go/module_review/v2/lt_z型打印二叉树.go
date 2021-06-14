/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/13 8:27 上午
# @File : lt_z型打印二叉树.go
# @Description :
# @Attention :
*/
package v2

func zigzagLevelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	r := make([][]int, 0)
	toogle := true
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
			if toogle {
				reverseList222(list)
			}
		}
		r = append(r, list)
		toogle = !toogle
	}
	return r
}
func reverseList222(r []int) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}
