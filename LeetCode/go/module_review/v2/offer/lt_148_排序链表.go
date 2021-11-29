/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/16 8:58 下午
# @File : lt_148_排序链表.go
# @Description :
# @Attention :
*/
package offer

// 关键:
// 归并排序法 1. 先找到中点,然后找到左边和右边 2 最后再进行归并
func sortList(head *ListNode) *ListNode {
	return sortListMerge1(head)
}

func sortListMerge1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := sortListMerge2FindMid(head)
	tail := mid.Next
	mid.Next = nil

	left := sortListMerge1(head)
	right := sortListMerge1(tail)

	return sortListMerge3(left, right)
}
func sortListMerge2FindMid(node *ListNode) *ListNode {
	fast := node.Next
	slow := node
	for nil != fast.Next && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func sortListMerge3(l1, l2 *ListNode) *ListNode {
	dumy := &ListNode{}
	tmp := dumy
	for nil != l1 && nil != l2 {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if nil != l1 {
		tmp.Next = l1
	}
	if nil != l2 {
		tmp.Next = l2
	}
	return dumy.Next
}
