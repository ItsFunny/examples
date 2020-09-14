/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-21 09:21 
# @File : _82_Remove_Duplicates_from_Sorted_List_II.go
# @Description : 有重复的元素就直接全部删除了,可能存在删除头节点的情况,所以需要引入dummyNode
  并且是重复的都不能存在,所以需要 知道 当前节点,next 节点和 next.next节点
# @Attention :
	唯一需要注意的就是小心头结点会被删除
*/
package list

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	walkerNode := dummy
	var rmVal int
	for nil != walkerNode && nil != walkerNode.Next && nil != walkerNode.Next.Next {
		if walkerNode.Next.Val == walkerNode.Next.Next.Val {
			// 删除节点
			rmVal = walkerNode.Next.Val
			for  nil != walkerNode.Next && walkerNode.Next.Val == rmVal {
				walkerNode.Next = walkerNode.Next.Next
			}
		} else {
			// 移动元素
			walkerNode = walkerNode.Next
		}
	}

	return dummy.Next
}
