/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-17 08:44 
# @File : lt_817_Linked_List_Components.go
# @Description : 
# @Attention : 
*/
package list

/*
	可以认为是 找公共字符串组合
	解答方法很简单,核心就是 连接(表明下一个)=>找尾节点" 如果这个值在 G中,并且 节点的下一个值不在G中,则表明是其中一个组件,则++
 */
func numComponents(head *ListNode, G []int) int {
	if G == nil {
		return 0
	}
	set := make(map[int]struct{}, len(G))
	for i := 0; i < len(G); i++ {
		set[G[i]] = struct{}{}
	}
	count := 0
	for nil != head {
		if _, exist := set[head.Val]; !exist {
			head = head.Next
			continue
		}
		if head.Next == nil {
			count++
		} else {
			if _, exist := set[head.Next.Val]; !exist {
				count++
			}
		}
		head = head.Next
	}
	return count
}
