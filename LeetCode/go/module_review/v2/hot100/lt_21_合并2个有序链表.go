/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/14 10:31 上午
# @File : lt_21_合并2个有序链表.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 直接暴力遍历即可
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy
	for nil != l1 && nil != l2 {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if nil != l1 {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}
	return dummy.Next
}
