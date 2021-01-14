/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-27 09:53 
# @File : lt_142_Linked_List_Cycle_II.go
# @Description :    如果有环,返回入环节点
# @Attention :
	如果碰撞了,则fast 到head 处步调一致移动
	错误点:
		当碰撞之后,slow 要移动到slow的下一个节点 既slow=slow.next
*/
package list

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head.Next
	slow := head
	for nil != fast && nil != fast.Next {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			slow = slow.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}

	return nil
}
