/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-23 19:34 
# @File : lt_1669_Merge_In_Between_Linked_Lists.go
# @Description : 
# @Attention : 
*/
package list

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := &ListNode{Val: 0, Next: list1,}
	cur := 0
	tmp, prev := dummy, dummy
	var  bNode *ListNode
	for ; nil != tmp; {
		cur++
		if cur == a {
			prev = tmp
			// aNode = tmp
		} else if cur == b {
			bNode = tmp
		}
		tmp = tmp.Next
	}
	prev.Next = list2
	for tmp = list2; nil != tmp.Next; tmp = tmp.Next {
	}
	tmp.Next = bNode

	return dummy.Next
}
