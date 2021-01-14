/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-28 09:07 
# @File : lt_234_Palindrome_Linked_List.go
# @Description : 判断是否是回文字符串
# @Attention :
	找到中间节点,反转匹配即可
*/
package list

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	mid := findMid(head)
	tail := mid.Next
	mid.Next = nil
	tail = reverse(tail)

	for nil != head && nil != tail {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	for nil != node {
		temp := node.Next
		node.Next = prev

		prev = node
		node = temp
	}
	return prev
}

func findMid(node *ListNode) *ListNode {
	fast := node.Next
	slow := node
	for nil != fast && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
