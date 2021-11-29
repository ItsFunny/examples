/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/25 8:55 下午
# @File : lt_206_反转链表.go
# @Description :
# @Attention :
*/
package v2

// 关键: 初始化一个prev 节点, 以及很关键的,链表的操作都是,先连后断
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	for nil != head {
		next := head.Next
		head.Next = prev

		prev = head
		head = next
	}

	return prev
}
