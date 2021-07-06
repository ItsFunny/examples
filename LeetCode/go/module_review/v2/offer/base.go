/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-21 09:06
# @File : base.go
# @Description :
# @Attention :
*/
package offer

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	tmp := l
	for nil != tmp {
		s += strconv.Itoa(tmp.Val) + "->"
		tmp = tmp.Next
	}
	return s
}

func CreateNodeBy(values ...int) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	t := dummy
	for i := 0; i < len(values); i++ {
		v := &ListNode{Val: values[i]}
		t.Next = v
		t = t.Next
	}
	return dummy.Next
}
func CreateListNode(size int) *ListNode {
	n := &ListNode{Val: 0}
	dummy := &ListNode{
		Val:  0,
		Next: n,
	}
	for i := 1; i <= size; i++ {
		v := &ListNode{Val: i}
		n.Next = v
		n = v
	}
	return dummy.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Prev  *Node
	Child *Node
	Next  *Node
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}
