/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 17:07 
# @File : of_剑指_Offer_32_II_从上到下打印二叉树_II.go
# @Description : 
# @Attention : 
*/
package offer

func levelOrder2(root *TreeNode) [][]int {
	return bbfs(root)
}
func bbfs(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		level := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			tempNode := queue[0]
			if nil != tempNode.Left {
				queue = append(queue, tempNode.Left)
			}
			if nil != tempNode.Right {
				queue = append(queue, tempNode.Right)
			}
			level = append(level, tempNode.Val)
			queue = queue[1:]
		}
		result = append(result, level)
	}
	return result
}
