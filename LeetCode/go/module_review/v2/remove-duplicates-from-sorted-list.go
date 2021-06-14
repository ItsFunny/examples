/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/25 9:50 上午
# @File : remove-duplicates-from-sorted-list.go
# @Description :	给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
  既 要用当前节点与后面的所有节点匹配
# @Attention :
*/
package v2


func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for ; nil != node; node = node.Next {
		// 拿当前节点与后面的所有节点匹配
		for ; node.Next != nil && node.Next.Val == node.Val; {
			node.Next = node.Next.Next
		}
	}
	return head
}
