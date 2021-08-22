/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/19 8:25 上午
# @File : jz_38_二叉树的深度.go
# @Description :
# @Attention :
*/
package offer

func TreeDepth(pRoot *TreeNode) int {
	return TreeDepthLevelTree(pRoot)
}
func TreeDepthLevelTree(pRoot *TreeNode) int {
	if pRoot==nil{
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, pRoot)
	r := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if nil != node.Left {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
		}
		r++
	}
	return r
}
