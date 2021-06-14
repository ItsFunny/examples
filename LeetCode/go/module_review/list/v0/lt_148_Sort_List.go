/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-27 09:03 
# @File : lt_148_Sort_List.go
# @Description :   链表排序
# @Attention :
	类似于快速排序
	错误的地方:
		1. mergeSort的返回值判断不仅要判断head还要判断head.next,并且返回值是head
		2. 需要使用dummy node
*/
package v0

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if nil == head || head.Next == nil {
		return head
	}
	midNode := midNode(head)
	tail := midNode.Next
	midNode.Next = nil
	left := mergeSort(head)
	right := mergeSort(tail)
	res := mergeTwList(left, right)
	return res
}
func mergeTwList(l1, l2 *ListNode) *ListNode {
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
func midNode(node *ListNode) *ListNode {
	if nil == node {
		return nil
	}
	fast := node.Next
	slow := node
	for nil != fast && nil != fast.Next {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
