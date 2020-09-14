/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-21 09:05 
# @File : _83_Remove_Duplicates_from_Sorted_List.go
# @Description : 删除重复的元素
# @Attention :
 遍历元素,一个一个匹配删除即可
*/
package list
//
// func deleteDuplicates(head *ListNode) *ListNode {
// 	walkerNode := head
// 	for nil != walkerNode {
// 		nextNode := walkerNode.Next
// 		if nil != nextNode && nextNode.Val == walkerNode.Val {
// 			// 重复元素,删除
// 			walkerNode.Next = walkerNode.Next.Next
// 		} else {
// 			walkerNode = walkerNode.Next
// 		}
// 	}
//
// 	return head
// }
