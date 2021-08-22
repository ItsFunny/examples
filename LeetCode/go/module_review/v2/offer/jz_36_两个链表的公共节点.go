/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/16 8:54 上午
# @File : jz_36_两个链表的公共节点.go
# @Description :
# @Attention :
*/
package offer

// 关键 p1+p2=p2+p1
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	t1, t2 := pHead1, pHead2
	for t1 != t2 {
		if t1 != nil {
			t1 = t1.Next
		} else {
			t1 = pHead2
		}
		if nil != t2 {
			t2 = t2.Next
		} else {
			t2 = pHead1
		}
	}
	return t1
}
