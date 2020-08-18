/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 15:21 
# @File : of_剑指_Offer_25_合并两个排序的链表.go
# @Description : 有序链表排序
# @Attention : 
*/
package offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	walerNode := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			walerNode.Next = l1
			l1 = l1.Next
		} else {
			walerNode.Next = l2
			l2 = l2.Next
		}
		walerNode = walerNode.Next
	}
	if nil != l1 {
		walerNode.Next = l1
		walerNode = walerNode.Next
	}
	if nil != l2 {
		walerNode.Next = l2
	}
	return dummy.Next
}
