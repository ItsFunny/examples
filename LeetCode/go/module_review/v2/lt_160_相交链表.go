/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/21 1:11 下午
# @File : lt_160_相交链表.go
# @Description :
# @Attention :
*/
package v2

// 找到两个单链表相交的节点
// 关键:
// 双指针: pa,pb ,如果pa为空了,则 pa 移动到headB重新开始,pb同理
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
