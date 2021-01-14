/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 16:58 
# @File : of_剑指_Offer_32_I_从上到下打印二叉树.go
# @Description : 
# @Attention : 
*/
package offer

func levelOrder(root *TreeNode) []int {
	return bfs(root)
}
func bfs(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([]int, 0)
	for len(queue) > 0 {
		tempNode := queue[0]
		result = append(result, tempNode.Val)
		if nil != tempNode.Left {
			queue = append(queue, tempNode.Left)
		}
		if nil != tempNode.Right {
			queue = append(queue, tempNode.Right)
		}
		queue = queue[1:]
	}
	return result
}
