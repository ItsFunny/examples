/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-14 09:21 
# @File : lt_237_Delete_Node_in_a_Linked_List.go
# @Description : 
# @Attention : 
*/
package v0


/*
	毫无意义的一道题目
 */
func deleteNode(node *ListNode) {
	newValue := node.Next.Val
	node.Next = node.Next.Next
	node.Val = newValue
}