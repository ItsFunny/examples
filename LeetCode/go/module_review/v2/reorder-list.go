/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/29 2:34 下午
# @File : reorder-list.go
# @Description :

给定一个单链表L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
# @Attention :
*/
package v2

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 1. 找到中点
	mid := reorderListFindMiddle(head)
	// 2. 断开mid
	next := mid.Next
	mid.Next = nil
	// 3. 反转后半段
	afterReverse := reorderListReverse(next)
	// 4. 拼接2端
	head = reorderListMergeList(head, afterReverse)
}
func reorderListMergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	move := dummy
	toogle := true
	for nil != l1 && nil != l2 {
		if toogle {
			move.Next = l1
			l1 = l1.Next
		} else {
			move.Next = l2
			l2 = l2.Next
		}
		toogle = !toogle
		move = move.Next
	}

	if nil != l1 {
		move.Next = l1
	}
	if nil != l2 {
		move.Next = l2
	}

	return dummy.Next
}
func reorderListFindMiddle(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for nil != fast.Next && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reorderListReverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for nil != cur {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev
}
