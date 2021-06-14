/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/25 11:23 上午
# @File : reverse-linked-list.go
# @Description : 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
# @Attention :
*/
package v2

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for nil != cur {
		tNode := cur.Next
		cur.Next = prev
		prev = cur
		cur = tNode
	}
	return prev
}
