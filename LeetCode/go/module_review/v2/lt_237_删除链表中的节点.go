/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/29 9:13 上午
# @File : lt_237_删除链表中的节点.go
# @Description :
# @Attention :
*/
package v2

// 关键: 直接赋值即可
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
