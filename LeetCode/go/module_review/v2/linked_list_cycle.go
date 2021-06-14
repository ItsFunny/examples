/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/29 4:46 下午
# @File : linked_list_cycle.go
# @Description :
给定一个链表，判断链表中是否有环。

如果链表中有某个节点，
可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
# @Attention :
*/
package v2

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head.Next
	slow := head
	for fast != nil && nil != fast.Next {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
