/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/8 9:38 上午
# @File : jz_26_二叉树和双向链表.go
# @Description :
# @Attention :
*/
package offer

func Convert(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}
	var prev *TreeNode
	var r *TreeNode
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || nil != root {
		for nil != root {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root.Left = prev
		if nil != prev {
			prev.Right = root
		}
		if nil == r {
			r = root
		}
		prev = root
		// if nil != root.Right {
		// 	root=root.Right
		// }
		// 注意,这一步中,不能有判断right 是否为空?
		// 为什么? 原因: 当root.Right为空的时候,则此时root 依然为原先的值,则会进入上层的for循环,导致数据重复再来一遍
		// 所以: root 要么是有值(right有值),否则必须设置为空
		root = root.Right
	}

	return r
}
func convertInOrdererTree(node *TreeNode) {

}
