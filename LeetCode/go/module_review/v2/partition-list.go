/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/28 9:25 上午
# @File : partition-list.go
# @Description :
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

# @Attention :
*/
package v2

// 思路就是 >= x 的放在另外一个链表,然后最后相连
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	headDummy := &ListNode{Next: head}
	biggerDummmy := &ListNode{}
	tBigger := biggerDummmy
	head = headDummy
	for nil != head.Next {
		if head.Next.Val < x {
			head = head.Next
		} else {
			// 断开该节点
			t := head.Next
			head.Next = head.Next.Next
			tBigger.Next = t
			tBigger = tBigger.Next
		}
	}
	tBigger.Next=nil
	head.Next = biggerDummmy.Next
	return headDummy.Next
}
