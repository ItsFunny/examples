/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 10:25 
# @File : of_剑指_Offer_06_从尾到头打印链表.go
# @Description : 
# @Attention : 
*/
package offer

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	result := make([]int, 0)
	for nil != head {
		result = append(result, head.Val)
		head = head.Next
	}
	reverse(result)
	return result
}
func reverse(data []int) {
	for i, j := 0, len(data)-1; i <= j; {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
