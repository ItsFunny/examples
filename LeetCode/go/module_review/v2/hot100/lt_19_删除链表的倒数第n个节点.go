/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/13 8:56 下午
# @File : lt_19_删除链表的倒数第n个节点.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 快慢指针,快指针走n步
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := head, dummy
	// 关键:
	// 第一步: fast从head触发,快指针先走n步,slow 从dummy出发,当走到尾的时候,代表着中点
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 第二步: 慢指针从dummy出发,当fast走到nil之后,代表着slow 走到了第n步
	for nil != fast {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
