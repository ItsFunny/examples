/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/22 10:19 下午
# @File : lt_23_合并k个有序链表.go
# @Description :
# @Attention :
*/
package offer

func mergeKLists(lists []*ListNode) *ListNode {
	if nil == lists || len(lists) == 0 {
		return nil
	}
	mergeKListsWithM(lists, 1)
	return lists[0]
}

func mergeKListsWithM(lists []*ListNode, index int) {
	if len(lists) == index {
		return
	}

	first := lists[0]
	second := lists[index]
	dummy := &ListNode{}
	tmp := dummy
	for nil != first && nil != second {
		if first.Val < second.Val {
			tmp.Next = first
			first = first.Next
		} else {
			tmp.Next = second
			second = second.Next
		}
		tmp = tmp.Next
	}
	if nil != first {
		tmp.Next = first
	} else {
		tmp.Next = second
	}
	lists[0] = dummy.Next
	mergeKListsWithM(lists, index+1)
}
