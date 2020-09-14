/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-11 09:07 
# @File : lt_147_Insertion_Sort_List.go
# @Description : 
# @Attention : 
*/
package list

/*
	插入排序: 链表的实现方式
	1. for 循环放到数组中,再插入排序
	插入排序: 有序中插入
	2. for循环中,发现不匹配的,再从头开始
 */

func insertSort(data []int) {
	for i := 1; i < len(data); i++ {
		j := i - 1
		val := data[i]
		for ; j >= 0 && data[j] > val; j-- {
			data[j] = data[j+1]
		}
		data[j] = val
	}
}

func insertionSortList2(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	dummy := &ListNode{
		Next: head,
	}
	prev := head
	next := head.Next
	for nil != next {
		if next.Val < prev.Val {
			// 从头开始,然后插入
			walkerNode := dummy.Next
			for nil != walkerNode && walkerNode.Next != prev {
				walkerNode = walkerNode.Next
			}
			nNext := next.Next
			walkerNode.Next = next
			next.Next = prev
			prev.Next = nNext
			next = nNext
		} else {
			next = next.Next
			prev = prev.Next
		}
	}
	return dummy.Next
}

func insertionSortList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	walkerNode := head.Next
	prev := head
	var pprev *ListNode
	prev.Next = pprev
	for nil != walkerNode {
		val := walkerNode.Val
		for nil != prev && prev.Val > val {

		}
		pprev = prev
		prev = walkerNode
	}
	return head
}
