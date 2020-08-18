/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 15:02 
# @File : of_剑指 Offer_22_链表中倒数第k个节点.go
# @Description : 双指针
# @Attention : 
*/
package offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	dummyNode := ListNode{
		Val:  0,
		Next: head,
	}
	slow, fast := dummyNode.Next, dummyNode.Next
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for nil != fast {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
