/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 16:50 
# @File : lt_24_Swap_Nodes_in_Pairs.go
# @Description : 
# @Attention : 
*/
package recusive

func swapPairs(head *ListNode) *ListNode {
	return swip(head,false)
}

func swip(node *ListNode, swap bool)*ListNode{
	if node==nil || node.Next==nil{
		return node
	}

	next:=node.Next
	nnext:=next.Next
	next.Next=node
	node.Next=swip(nnext,!swap)
	return next
}