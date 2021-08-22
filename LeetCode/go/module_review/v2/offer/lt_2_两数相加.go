/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/30 8:45 上午
# @File : lt_2_两数相加.go
# @Description :
# @Attention :
*/
package offer

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		l1Val int
		l2Val int
	)
	var head *ListNode
	var tail *ListNode
	more := 0
	for nil != l1 || nil != l2 {
		if nil != l1 {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			l2Val = l2.Val
			l2 = l2.Next
		}
		value := l1Val + l2Val + more
		value, more = value%10, value/10
		if head == nil {
			head = &ListNode{Val: value}
			tail = head
		} else {
			tail.Next = &ListNode{Val: value}
			tail = tail.Next
		}

		l1Val = 0
		l2Val = 0
	}
	if more>0{
		tail.Next=&ListNode{Val: more}
	}

	return head
}
