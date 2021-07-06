/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/2 8:39 上午
# @File : jz_14_链表倒数k个节点.go
# @Description :
# @Attention :
*/
package offer

// 对于这种无法获得长度,但是求倒数的情况
// 都可以用双指针解决,1个先走k步,另外一个从头开始,前者到末尾,则后者刚好到倒数k步
func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
	// write code here
	first,second:=pHead,pHead
	for i:=0;i<k;i++{
		if first==nil{
			return nil
		}
		first=first.Next
	}
	for nil!=first{
		first=first.Next
		second=second.Next
	}
	return second
}
