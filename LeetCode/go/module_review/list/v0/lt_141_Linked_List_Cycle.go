/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-27 09:50 
# @File : lt_141_Linked_List_Cycle.go
# @Description : 判断是否有环
# @Attention :  快慢指针
*/
package v0

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next
	slow := head
	for nil != fast && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}

	}
	return false
}
