/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/28 9:08 下午
# @File : sort_list.go
# @Description :
给你链表的头结点head请将其按 升序 排列并返回 排序后的链表 。
你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
# @Attention :
*/
package v2

func sortList(head *ListNode) *ListNode {
	return sortListMergeSort(head)
}
func sortListMergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := sortListFindMiddle(head)
	tail := mid.Next
	mid.Next = nil

	left := sortListMergeSort(head)
	right := sortListMergeSort(tail)
	return sortListMergeList(left, right)
}
func sortListFindMiddle(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func sortListMergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if nil != l1 {
		temp.Next = l1
	}
	if nil != l2 {
		temp.Next = l2
	}

	return dummy.Next
}
