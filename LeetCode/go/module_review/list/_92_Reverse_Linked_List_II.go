/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-22 08:55 
# @File : _92_Reverse_Linked_List_II.go
# @Description :    反转下标 m-n 的链表
# @Attention :
	截取 m-n的链表  ,进行反转
	或者是
*/
package list

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 可能涉及到头节点,所以需要引入dummy node
	dummy := &ListNode{}
	dummy.Next = head
	var (
		prev        *ListNode
		reversePrev *ListNode
		reverseHead *ListNode
		next        *ListNode
	)
	i := 0
	for ; i < m; i++ {
		if i == m-1 {
			reverseHead = head
			reversePrev = prev
		} else {
			prev = head
			head = head.Next
		}
	}
	prev.Next = nil
	for j := i; j < n; j++ {
		temp := head.Next
		head.Next = next

		next = temp
		head = temp
	}

	reversePrev.Next = head
	head.Next = next

	return dummy.Next
}
func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	for nil != node {
		temp := node.Next
		node.Next = prev

		prev = node
		node = temp
	}

	return prev
}
