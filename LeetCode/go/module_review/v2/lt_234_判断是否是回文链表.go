/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/28 9:48 上午
# @File : lt_234_判断是否是回文链表.go
# @Description :
# @Attention :
*/
package v2

// 关键:
// 1. 找中点
// 2. 中点之后进行反转
// 3. 匹配判断
func isPalindrome2(head *ListNode) bool {
	mid := isPalindrome2FindMid(head)
	midNext := mid.Next
	mid.Next = nil

	midNext = isPalindrome2Reverse(midNext)

	for nil != head && nil != midNext {
		if head.Val != midNext.Val {
			return false
		}
		midNext = midNext.Next
		head = head.Next
	}
	return true
}

func isPalindrome2FindMid(node *ListNode) *ListNode {
	if nil == node {
		return nil
	}
	fast := node.Next
	slow := node
	for nil != fast && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func isPalindrome2Reverse(node *ListNode) *ListNode {
	var prev *ListNode
	cur := node
	for nil != cur {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}
