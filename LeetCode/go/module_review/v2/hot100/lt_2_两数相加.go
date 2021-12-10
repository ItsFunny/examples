/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/7 8:48 下午
# @File : lt_2_两数相加.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 2个节点,一个作为返回值,一个作为后续末尾值,因为最后可能会大于10,所以需要手动的添加一个值
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	needMore := false
	var ret *ListNode
	var tail *ListNode
	for nil != l1 || nil != l2 {
		tmp := 0
		if nil != l1 {
			tmp += l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			tmp += l2.Val
			l2 = l2.Next
		}
		if needMore {
			tmp++
		}
		if tmp >= 10 {
			needMore = true
			tmp %= 10
		} else {
			needMore = false
		}
		if nil == ret {
			ret = &ListNode{Val: tmp}
			tail = ret
		} else {
			tail.Next = &ListNode{Val: tmp}
			tail = tail.Next
		}
	}
	// 如果>10 ,还需要的是,将其末尾追加一个值
	if needMore {
		h := &ListNode{Val: 1}
		tail.Next = h
	}

	return ret
}
