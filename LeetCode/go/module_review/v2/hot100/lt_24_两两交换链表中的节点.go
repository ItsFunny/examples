/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/16 9:41 上午
# @File : lt_24_两两交换链表中的节点.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 使用dummy节点,用dummy来移动整个链表
// 每次交换都是交换dummy后的两个节点即可
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	tmp := dummy
	// tmp->1->2 =>  tmp->2->1
	for nil != tmp && tmp.Next != nil && tmp.Next.Next != nil {
		node1 := tmp.Next
		node2 := node1.Next
		tmp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		tmp = node1
	}
	return dummy.Next
}
