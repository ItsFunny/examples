/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 10:38 
# @File : _19_Remove_Nth_Node_From_End_of_List.go
# @Description : Given linked list: 1->2->3->4->5, and n = 2.
				 After removing the second node from the end, the linked list becomes 1->2->3->5.
				 移除链表中的元素并且返回头结点,要求O(n)
# @Attention : 
*/
package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}