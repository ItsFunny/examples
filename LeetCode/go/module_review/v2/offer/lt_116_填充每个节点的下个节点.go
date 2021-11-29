/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/25 10:32 下午
# @File : lt_116_填充每个节点的下个节点.go
# @Description :
# @Attention :
*/
package offer

// 解题关键:
// 题目的意思是: 层序遍历,然后每一层的元素,指向这层的下一个元素
// 所以关键是层序遍历 BFS: queue
func connect(root *LNode) *LNode {
	if nil == root {
		return nil
	}
	queue := make([]*LNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		// 然后遍历这层,建立连接关系
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if nil != node.Left {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
