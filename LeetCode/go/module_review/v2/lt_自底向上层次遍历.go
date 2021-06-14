/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/13 8:05 上午
# @File : lt_自底向上层次遍历.go
# @Description :
# @Attention :
*/
package v2

func levelOrderBottom(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	order11 := levelOrder11(root)
	reverse(order11)
	return order11
}
func reverse(r [][]int) {
	for i, j := 0, len(r)-1; i< j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
}
func levelOrder11(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	r := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		list := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if nil != node.Left {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
		}
		r = append(r, list)
	}
	return r
}
