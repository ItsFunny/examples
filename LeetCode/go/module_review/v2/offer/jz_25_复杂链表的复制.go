/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/8 9:14 上午
# @File : jz_25_复杂链表的复制.go
# @Description :
# @Attention :
*/
package offer

func Clone(head *RandomListNode) *RandomListNode {
	// write your code here
	if nil == head {
		return nil
	}
	m := make(map[*RandomListNode]*RandomListNode)
	temp := head
	for nil != temp {
		m[temp] = &RandomListNode{Label: temp.Label}
		temp = temp.Next
	}
	r := &RandomListNode{Label: head.Label}
	temp = r
	for nil != head {
		temp.Next = m[head.Next]
		temp.Random = m[head.Random]

		temp = temp.Next
		head = head.Next
	}

	return r
}

func clone(node *RandomListNode, m map[*RandomListNode]*RandomListNode) *RandomListNode {
	if nil == node {
		return nil
	}
	var r *RandomListNode
	v, exist := m[node]
	if exist {
		r = v
	} else {
		r = &RandomListNode{
			Label: node.Label,
		}
		m[node] = r
	}
	r.Next = clone(node.Next, m)
	r.Random = clone(node.Random, m)
	return r
}
