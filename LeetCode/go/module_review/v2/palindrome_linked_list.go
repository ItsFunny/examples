/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/29 10:46 下午
# @File : palindrome_linked_list.go
# @Description :
判断一个链表是否为回文链表。
# @Attention :
*/
package v2

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	mid := isPalindromeFindMiddle(head)
	// 断开
	next := mid.Next
	mid.Next = nil
	rev := isPalindromeReverse(next)
	for nil != head && nil != rev {
		if head.Val != rev.Val {
			return false
		}
		head = head.Next
		rev = rev.Next
	}
	// return  rev == nil && head==nil
	return true
}
func isPalindromeReverse(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for nil != cur {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
func isPalindromeFindMiddle(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for nil != fast && nil != fast.Next {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
