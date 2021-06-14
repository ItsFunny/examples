/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/25 10:08 上午
# @File : remove-duplicates-from-sorted-list-ii.go
# @Description :
存在一个按升序排列的链表，给你这个链表的头节点 head ，
请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
返回同样按升序排列的结果链表。
# @Attention :
可能删除头节点
有序,所以如果相邻的不相同,之后的必然不相同
*/
package v2


func deleteDuplicates2(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	head = dummy
	var rmValue int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmValue = head.Next.Val
			// 删除之后的所有跟这个相同的元素
			for ; nil != head.Next && head.Next.Val == rmValue; {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}

	return dummy.Next
}
