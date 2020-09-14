/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-11 08:59 
# @File : main.go
# @Description : 
# @Attention : 
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	r := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	ReorderList(r)
	fmt.Println(r)
}
