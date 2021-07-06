/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/26 4:07 下午
# @File : lt_offer_链表反转.go
# @Description :
# @Attention :
*/
package v2

// 链表反转

func reversePrint(head *ListNode) []int {
	if nil == head {
		return nil
	}
	stack := make([]int, 0)
	for nil != head {
		stack = append(stack, head.Val)
		head = head.Next
	}
	r := make([]int, 0)
	for len(stack) > 0 {
		r = append(r, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return r
}
