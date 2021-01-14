/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-14 09:13 
# @File : base.go
# @Description :
 链表核心知识点:
 1. nil退出条件处理
 2. dummy node 哑巴节点
 3. 快慢指针
	3.1 找到链表的中间节点
 4. 链表插入
 5. 链表删除
 6. 反转链表
 7. 合并链表
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
			walkerNode = walkerNode.Next.Next
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
		head.Next = prev
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

// 在  O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
// 思路: 模拟快排的思路,找到中间点排序
// 重点:
// 1. 快慢指针 获取中间元素
// 2. mergeSort:归并排序(简易快排)的时候断开中间节点
// 3. 分治法的退出条件是 为空,或者是 只有一个节点
func sortList(head *ListNode) *ListNode {
	return qSortLinkedList(head)
}
func qSortLinkedList(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	middle := paration(node)
	// 排序之前需要断开,因为链表不像数组,链表是有前后连接的
	after := middle.Next
	middle.Next = nil
	// 对左右进行排序
	left := qSortLinkedList(node)
	right := qSortLinkedList(after)
	result := mergeLists(left, right)
	return result
}
func mergeLists(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	walkerNode := dummy
	for nil != left && right != nil {
		if left.Val < right.Val {
			walkerNode.Next = left
			left = left.Next
		} else {
			walkerNode.Next = right
			right = right.Next
		}
		walkerNode = walkerNode.Next
	}

	for nil != left {
		walkerNode.Next = left
		left = left.Next
		walkerNode = walkerNode.Next
	}
	for nil != right {
		walkerNode.Next = right
		right = right.Next
		walkerNode = walkerNode.Next
	}
	return dummy.Next
}

// 链表找中间元素的方法: 就是通过快慢指针,当快指针到末尾的时候,慢指针刚好到中间
func paration(node *ListNode) *ListNode {
	fast := node.Next
	slow := node
	for nil != fast && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 判断是否有环
// 快慢指针 如果有环每次前进都会使得距离缩短-1
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil || head.Next.Next == nil {
		return true
	}
	fast := head.Next.Next
	slow := head.Next
	for nil != fast && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 给定一个链表，返回链表开始入环的第一个节点。  如果链表无环，则返回  null。
// 思路: 快慢指针,相遇代表有环,有环之后,慢指针到头部,两者步调一致,相遇的既为入环的节点
func detectCycleNode(head *ListNode) *ListNode {
	if nil == head || head.Next == nil {
		return nil
	}
	fast := head.Next
	slow := head
	for nil != fast && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return slow
		}
	}

	return nil
}

// 判断链表是否是一个回文链表

func isPalindrome(head *ListNode) bool {
	if nil == head {
		return true
	}
	// 找到中间节点
	fast := head.Next
	slow := head
	for nil != fast && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow
	// 断开连接
	after := mid.Next
	mid.Next = nil
	// 反转后半段链表
	after = revSe(after)
	// 判断
	for nil != head && nil != after {
		if head.Val != after.Val {
			return false
		}
		head = head.Next
		after = after.Next
	}
	return true
}

func revSe(node *ListNode) *ListNode {
	var prev *ListNode
	for nil != node {
		tempNode := node.Next
		tempNode.Next = prev

		prev = node
		node = tempNode
	}
	return prev
}

// 给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。 要求返回这个链表的 深拷贝
