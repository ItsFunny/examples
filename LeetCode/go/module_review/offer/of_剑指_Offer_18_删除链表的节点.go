/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 14:41 
# @File : of_剑指_Offer_18_删除链表的节点.go
# @Description : 
# @Attention : 
*/
package offer

func deleteNode(head *ListNode, val int) *ListNode {
	if nil == head {
		return nil
	}
	dummyNode := &ListNode{
		Next: head,
	}
	walkerNode := dummyNode
	prevNode := walkerNode
	for nil != walkerNode {
		if walkerNode.Val == val {
			break
		}
		prevNode = walkerNode
		walkerNode = walkerNode.Next
	}
	prevNode.Next = walkerNode.Next
	return dummyNode.Next
}
