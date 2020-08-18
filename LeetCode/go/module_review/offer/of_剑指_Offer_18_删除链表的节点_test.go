/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 14:41
# @File : of_剑指_Offer_18_删除链表的节点.go
# @Description :
# @Attention :
*/
package offer

import (
	"testing"
)

func Test_deleteNode(t *testing.T) {
	deleteNode(&ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}, 5)
}
