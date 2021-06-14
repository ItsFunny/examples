/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-18 09:06 
# @File : lt_1367_Linked_List_in_Binary_Tree.go
# @Description : 
# @Attention : 
*/
package v0

/*
	判断链表中的节点是否在树中
 */

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root==nil{
		return false
	}
	return isSubPathDfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
func isSubPathDfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return isSubPathDfs(head.Next, root.Left) || isSubPathDfs(head.Next, root.Right)
}
