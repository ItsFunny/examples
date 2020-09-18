/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-18 08:55 
# @File : lt_1171_Remove_Zero_Sum_Consecutive_Nodes_from_Linked_List.go
# @Description : 
# @Attention : 
*/
package list

/*
	去除和抵消为0的值
	关键: 对每个元素的下标取和,如果 有重复的说明,[x,y] 之间的是可以抵消掉的
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	dummy := &ListNode{
		Next: head,
	}
	flagMap := make(map[int]*ListNode, 0)
	sum := 0
	for node := dummy; nil != node; node = node.Next {
		sum += node.Val
		flagMap[sum] = node
	}

	sum = 0
	for node := dummy; nil != node; node = node.Next {
		sum += node.Val
		node.Next = flagMap[sum].Next
	}
	return dummy.Next
}
