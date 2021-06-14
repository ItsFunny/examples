/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-10 08:37 
# @File : lt_109_Convert_Sorted_List_to_Binary_Search_Tree.go
# @Description : 
# @Attention : 
*/
package v0

/*
	有序列表,构建高度差小于1的树
	1. 既然是有序,则中间构建,左孩子都是小于它的,右孩子都是大于它的
	2. 快慢指针找中间节点->分治法
	3. 注意分治法的退出条件
 */

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func buildTree(left *ListNode, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := findMiddleNode(left, right)
	root := &TreeNode{Val: mid.Val,}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}

func findMiddleNode(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
