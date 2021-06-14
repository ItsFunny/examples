/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/29 4:53 下午
# @File : linked_list_cycle_ii.go
# @Description :
给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回null。

为了表示给定链表中的环，
我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

进阶：

你是否可以使用 O(1) 空间解决此题？

# @Attention :
*/
package v2

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast := head.Next
	slow := head
	for nil != fast && fast.Next != nil {
		if fast == slow {
			fast=head
			slow=slow.Next
			for nil != slow {
				if fast == slow {
					return slow
				}
				slow = slow.Next
				fast = fast.Next
			}
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}
