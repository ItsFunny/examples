/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-21 09:41 
# @File : _206_Reverse_Linked_List.go
# @Description : 反转链表,涉及到头结点,所以需要引入dummyNode
# @Attention :
	返回值,因为是反转,所以需要返回的是prev的值

*/
package v0

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var temp *ListNode
	for nil != head {
		temp = head.Next
		head.Next = prev

		prev=head
		head=temp
	}

	return prev
}
