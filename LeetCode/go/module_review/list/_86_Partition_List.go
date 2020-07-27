/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-23 09:21 
# @File : _86_Partition_List.go
# @Description : 既然 给定一个数字,小于的全在左边,大于的全在右边
# @Attention :
 第一种方法是链表排序,然后将值插入
 第二种方法是挑选出 大于或者小于的所有节点 ,如挑出大于的所有节点,再连接
*/
package list
func partition(head *ListNode, x int) *ListNode {
	// 思路：将大于x的节点，放到另外一个链表，最后连接这两个链表
	// check
	if head == nil {
		return head
	}
	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	tail := tailDummy
	headDummy.Next = head
	// 这里将head 从新移动到headDummy的缘故在于,for循环中是通过next判断的
	head = headDummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			// 移除<x节点
			t := head.Next
			head.Next = head.Next.Next
			// 放到另外一个链表
			tail.Next = t
			tail = tail.Next
		}
	}
	// 拼接两个链表
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}