/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-14 09:13 
# @File : base.go
# @Description : 
# @Attention : 
*/
package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
func deleteDuplicateNode(root *ListNode) {
	walkerNode := root
	for walkerNode != nil {
		for walkerNode.Next != nil && walkerNode.Val == walkerNode.Next.Val {
			walkerNode = walkerNode.Next
		}
		walkerNode = walkerNode.Next
	}
}

// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中   没有重复出现的数字。
func deleteDuplicate2(root *ListNode) *ListNode {
	if nil == root {
		return root
	}
	dummy := &ListNode{
		Val:  0,
		Next: root,
	}
	root = dummy
	var rmVal int
	for nil != root.Next && root.Next.Next != nil {
		// 说明下一个节点和下下一个节点的值相同,则需要删除这2个节点
		if root.Next.Val == root.Next.Next.Val {
			rmVal = root.Next.Val
			for root.Next != nil && root.Next.Val == rmVal {
				// 删除该节点
				root.Next = root.Next.Next
			}
		} else {
			// 节点发生移动
			root = root.Next
		}
	}
	return root
}

// 反转单链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for nil != head {
		temp := head.Next
		temp.Next = prev
		// prev 前移动 反转意味着 在前面的会变为在后面,所以head 会变为过去式
		prev = head
		// head 前移动 而之前的下一个节点会是先的next ,所以会变为新的head
		head = temp
	}
	return prev
}

// 反转单链表2
// 反转从位置  m  到  n  的链表。请使用一趟扫描完成反转。
// 思路：先遍历到 m 处，翻转，再拼接后续
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var prev *ListNode
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	prev = dummy
	for i := 0; i < m-1; i++ {
		prev = head.Next
	}
	// 当前prev 为m所处的上一个位置
	// 反转m-n 之间的节点
	node := reverse(prev, m-n)
	prev.Next = node
	return prev
}
func reverse(head *ListNode, count int) *ListNode {
	var prev *ListNode
	for i := 0; i < count && head != nil; i++ {
		temp := prev.Next
		prev.Next = prev

		prev = head
		head = temp
	}

	return prev
}

// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 通过dummyNode 实现
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	// 临时节点
	tempNode := dummy
	for nil != l1 && nil != l2 {
		if l1.Val < l2.Val {
			tempNode.Next = l1
			l1 = l1.Next
		} else {
			tempNode.Next = l2
			l2 = l2.Next
		}
		tempNode = tempNode.Next
	}

	// 可能l1 更长
	for nil != l1 {
		tempNode.Next = l1
		l1 = l1.Next
		tempNode = tempNode.Next
	}

	// 可能l2 更长
	for nil != l2 {
		tempNode.Next = l2
		l2 = l2.Next
		tempNode = tempNode.Next
	}

	return dummy.Next
}

// 分割链表
// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于  x  的节点都在大于或等于  x  的节点之前。
// 链表合并问题
// 思路: 找到该节点 ,同时将大于该节点的 放到另外的节点,最后再合并
func partition(head *ListNode, x int) *ListNode {
	var concreteNode *ListNode
	dummyLagerNode := &ListNode{}
	dummyWalkerNode := dummyLagerNode
	headDummy := &ListNode{}
	walkerNode := head
	for nil != walkerNode && nil != walkerNode.Next {
		if walkerNode.Val == x {
			concreteNode = walkerNode
			walkerNode.Next = walkerNode.Next.Next
		} else if walkerNode.Val > x {
			dummyWalkerNode.Next = walkerNode
			dummyWalkerNode = dummyWalkerNode.Next
		} else {
			if headDummy.Next == nil {
				headDummy.Next = walkerNode
			}
			walkerNode.Next = walkerNode.Next.Next
		}
		walkerNode = walkerNode.Next
	}

	// 拼接
	walkerNode.Next, concreteNode.Next = concreteNode, dummyLagerNode.Next

	return headDummy.Next
}
