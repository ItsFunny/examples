/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-27 09:20 
# @File : lt_143_Reorder_List.go
# @Description : 
# @Attention : 
*/
package v0

/*
	对链表重排序
	最简单的方式就是 依次从头部和尾部获取数据拼接
	1. 直接遍历存储下标即可
 */
func reorderList(head *ListNode) {
	nodes := make([]*ListNode, 0)
	for walkerNode := head; nil != walkerNode; {
		nodes = append(nodes, walkerNode)
		walkerNode = walkerNode.Next
	}
	dummy := &ListNode{}

	walkerNode := dummy.Next
	isHead := true
	for i, j := 0, len(nodes)-1; i <= j; {
		if isHead {
			walkerNode = nodes[i]
			i++
		} else {
			walkerNode = nodes[j]
			j--
		}
		walkerNode = walkerNode.Next
		isHead = !isHead
	}
	head = dummy.Next
}

// 找中点链表逆序之后,再拼接
func ReorderList(head *ListNode) {
	if head==nil{
		return
	}
	mid := reorderListFindMid(head, nil)
	// 断开连接
	right := mid.Next
	mid.Next = nil
	// 逆序right
	right = reorderListReverse(right)
	// 合并
	head = reorderListMerge(head, right)
}

func reorderListMerge(node *ListNode, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	walkerNode := dummy
	first := true
	for nil != node && nil != node2 {
		if first {
			walkerNode.Next = node
			node = node.Next
		} else {
			walkerNode.Next = node2
			node2 = node2.Next
		}
		walkerNode = walkerNode.Next
		first = !first
	}
	for nil != node {
		walkerNode.Next = node
		node = node.Next
		walkerNode = walkerNode.Next
	}
	for nil != node2 {
		walkerNode.Next = node2
		node2 = node2.Next
		walkerNode = walkerNode.Next
	}
	return dummy.Next
}

func reorderListReverse(node *ListNode) *ListNode {
	var prev *ListNode
	for nil != node {
		temp := node.Next
		node.Next = prev

		prev = node
		node = temp
	}
	return prev
}

func reorderListFindMid(left *ListNode, right *ListNode) *ListNode {
	slow, fast := left, left

	for right != fast && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
