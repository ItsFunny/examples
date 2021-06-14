/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-10 09:26 
# @File : _103_Binary_Tree_Zigzag_Level_Order_Traversal.go
# @Description : 依旧还是层次遍历,不同的地方在于隔一层需要反转,第一层不需要反转
# @Attention : 
*/
package v0

func zigzagLevelOrder(root *TreeNode) [][]int {
	return levelOrderTree(root)
}
func levelOrderTree(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	reverse := false
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
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if reverse {
			rev(list)
		}
		reverse = !reverse
		result = append(result, list)
	}
	return result
}
func rev(data []int) {
	for i, j := 0, len(data)-1; i < j; {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
