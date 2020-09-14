/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-11 09:07
# @File : lt_147_Insertion_Sort_List.go
# @Description :
# @Attention :
*/
package list

import (
	"fmt"
	"testing"
)

func Test_insertionSortList2(t *testing.T) {
	r := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	insertionSortList2(r)
	fmt.Println(r)
}
