/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-15 07:07 
# @File : lt_430_Flatten_a_Multilevel_Doubly_Linked_List.go
# @Description : 
# @Attention : 
*/
package list

/*
	扁平化多级双向链表
	依据图解: 就是认为为一个二叉树, child 为left 指针,next 为right 指针,结果为dfs,同时需要保存上一个指针
	因为并不是直接收集值,而是还需要指针连接
 */
func flatten(root *Node) *Node {
	if nil == root {
		return nil
	}
	dummy := &Node{
		Next: root,
	}
	cur, prev := dummy, dummy

	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		prev.Next = cur
		cur.Prev = prev
		if nil != cur.Next {
			stack = append(stack, cur.Next)
		}
		if nil != cur.Child {
			stack = append(stack, cur.Child)
			cur.Child = nil
		}
		prev = cur
	}

	dummy.Next.Prev = nil

	return dummy.Next
}
