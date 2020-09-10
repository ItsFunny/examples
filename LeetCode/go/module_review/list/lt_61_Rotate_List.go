/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-10 08:19 
# @File : lt_61_Rotate_List.go
# @Description : 
# @Attention : 
*/
package list

/*
	旋转链表
	将其构建成环
	找到新的头结点: 新的头节点在 n-(k%n) 处
	新的尾节点: 在 n-(k%n)-1 处
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if nil == head {
		return nil
	}
	walkerNode := head
	length := 1
	for nil != walkerNode.Next {
		walkerNode = walkerNode.Next
		length++
	}
	walkerNode.Next = head

	count := length - (k % length)
	walkerNode = head
	for i := 0; i < count-1; i++ {
		walkerNode = walkerNode.Next
	}
	tail := walkerNode
	newHead := tail.Next
	tail.Next = nil

	return newHead
}
