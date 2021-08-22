/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/21 11:02 上午
# @File : lt_19_删除链表倒数第n个节点.go
# @Description :
# @Attention :
*/
package offer

// 关键: 快慢指针,因为拿不到长度
// 快指针先走n步之后,当快指针刚好到末尾的时候,慢指针刚好到要删除的节点的前一个
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; nil != fast; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
