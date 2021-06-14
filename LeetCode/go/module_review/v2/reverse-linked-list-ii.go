/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/26 9:10 上午
# @File : reverse-linked-list-ii.go
# @Description :
给你单链表的头指针 head 和两个整数left 和 right ，
其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
# @Attention :
*/
package v2

// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	var (
// 		beforeReversePrev *ListNode
// 		prev              *ListNode
// 	)
//
// 	// dummy 用于返回 result
// 	dummy := &ListNode{
// 		Val:  0,
// 		Next: head,
// 	}
// 	i := 0
// 	// 1. 找到起始节点
// 	head = dummy
// 	for ; i < left; i++ {
// 		prev = head
// 		head = head.Next
// 		fmt.Println(head.String())
// 	}
// 	//  此时head处于m的下标处
// 	// 需要保存的是
// 	// 1. 当链表反转后的head: 使得 之前的 prev 可以连到这个新的head 既  beforeReversePrev.Next = newHead
// 	// 2. 链表反转前的head: 因为反转后原先的头节点,需要连接到 之前的尾节点的next
// 	beforeReversePrev = prev
// 	prev = nil
// 	cur := head
// 	oldHead:=head
// 	for j := i; j <= right; j++ {
// 		tmp := cur.Next
//
// 		cur.Next = prev
// 		prev = cur
// 		cur = tmp
// 	}
// 	beforeReversePrev.Next = prev
// 	oldHead.Next=cur
//
// 	return dummy.Next
// }

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	head=dummy
	// 3个节点
	// 1个是 反转前left 的prev
	// 第2个是反转前的head
	// 第3个是反转后的head

	var (
		headBeforeReverse *ListNode
	)

	i := 0
	for ; i < left; i++ {
		headBeforeReverse = head
		head = head.Next
	}

	// 反转后的head
	var prev *ListNode
	// 反转前的head
	oldHead := head
	for j := i; j <= right && oldHead != nil; j++ {
		tmp := head.Next

		head.Next = prev
		prev = head
		head = tmp
	}
	headBeforeReverse.Next = prev
	oldHead.Next = head

	return dummy.Next
}
