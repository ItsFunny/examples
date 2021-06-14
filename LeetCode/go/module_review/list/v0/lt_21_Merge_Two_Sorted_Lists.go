/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-23 09:11 
# @File : _21_Merge_Two_Sorted_Lists.go
# @Description : 两个有序链表合并为一个链表
# @Attention : 
*/
package v0

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	walkerNode := dummy
	for nil != l1 && nil != l2 {
		if l1.Val < l2.Val {
			walkerNode.Next = l1
			l1 = l1.Next
		} else {
			walkerNode.Next = l2
			l2 = l2.Next
		}
		walkerNode = walkerNode.Next
	}
	for nil != l1 {
		walkerNode.Next = l1
		l1 = l1.Next
		walkerNode = walkerNode.Next
	}
	for nil != l2 {
		walkerNode.Next = l2
		l2 = l2.Next
		walkerNode = walkerNode.Next
	}
	return dummy.Next
}
