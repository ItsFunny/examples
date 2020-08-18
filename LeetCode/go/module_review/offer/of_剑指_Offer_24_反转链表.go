/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 15:14 
# @File : of_剑指_Offer_24_反转链表.go
# @Description : 
# @Attention : 
*/
package offer

func reverseList(head *ListNode) *ListNode {
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
