/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-16 07:37 
# @File : lt_725_Split_Linked_List_in_Parts.go
# @Description : 
# @Attention : 
*/
package list

/*
	切割链表
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	nodeNum := 0
	temp:=root
	for temp != nil {
		nodeNum++
		temp = temp.Next
	}
	var result []*ListNode
	// 算出倍数以及余数 /  %
	width, remainder := nodeNum/k, nodeNum%k
	curr := root
	for i := 0; i < k; i++ { //每一组的数据长度
		head := curr
		// 难点在这，遍历width-1次，通过余数remainder进行判断是否全长度遍历
		for j := 0; j < (width-1+baseChange(i, remainder)) && curr != nil; j++ {
			curr = curr.Next
		}
		if curr != nil {
			curr.Next, curr = nil, curr.Next
		}
		result = append(result, head)

	}
	return result
}

func baseChange(min, max int) int {
	if min < max {
		return 1
	} else {
		return 0
	}
}