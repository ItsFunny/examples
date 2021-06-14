/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-14 09:25
# @File : lt_328_Odd_Even_Linked_List.go
# @Description :
# @Attention :
*/
package v0

import (
	"fmt"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	r := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	s := oddEvenList(r)
	fmt.Println(s)
}
