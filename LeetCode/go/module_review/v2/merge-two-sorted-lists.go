/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/28 9:05 上午
# @File : merge-two-sorted-lists.go
# @Description :
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

# @Attention :
*/
package v2

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		// head 移动从而进行匹配
		head = head.Next
	}
	// 为什么这列不能
	if nil != l1 {
		head.Next = l1
	}
	if nil != l2 {
		head.Next = l2
	}

	return dummy.Next
}
