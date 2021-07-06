/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/2 9:03 上午
# @File : jz_16_合并2个有序的链表.go
# @Description :
# @Attention :
*/
package offer

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	for nil != pHead1 && nil != pHead2 {
		if pHead1.Val < pHead2.Val {
			node.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			node.Next = pHead2
			pHead2 = pHead2.Next
		}
		node = node.Next
	}
	if nil != pHead1 {
		node.Next = pHead1
	}

	if nil != pHead2 {
		node.Next = pHead2
	}

	return dummy.Next
}
