/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/2 8:53 上午
# @File : jz_15_反转链表.go
# @Description :
# @Attention :
*/
package offer

func ReverseList(pHead *ListNode) *ListNode {
	var prev *ListNode
	for nil != pHead {
		next := pHead.Next

		pHead.Next=prev
		prev=pHead
		pHead=next
	}

	return prev
}
