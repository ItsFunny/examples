/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-17 09:25 
# @File : lt_1019_Next_Greater_Node_In_Linked_List.go
# @Description : 
# @Attention : 
*/
package list

/*
	1. 常见思路为 2个for 循环 O(n^2)的方式,但是怎么优化呢
	从题目可知,数据不是有序的,所以找中点不行,并且没必要,因为肯定得for循环一次
	双指针? 因为是链表,所以无法实现
	如何才能从O(n^2) 优化呢,要么优化成O(logn) 或者是优化成O(n)的方式 ,前者的话得转成树,后者的话
	可以遍历放到数组,空间换时间的方式

		结合栈: 栈用来存储
 */
func nextLargerNodes(head *ListNode) []int {
}
