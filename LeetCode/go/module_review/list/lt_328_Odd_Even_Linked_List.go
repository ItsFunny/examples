/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-14 09:25 
# @File : lt_328_Odd_Even_Linked_List.go
# @Description : 
# @Attention : 
*/
package list

/*
	奇偶链表
	将奇数节点放原来的节点,将偶数节点放在新的节点,最后再连接即可
	双指针即可,一个指向头部odd,一个指向头部的next代表偶数节点even,最后相连即可
 */

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	evenHead := head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func oddEvenList2(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	evenDummy := &ListNode{}
	evenWalker := evenDummy
	var prev *ListNode
	for i, walkerNode := 1, head; nil != walkerNode && nil != evenWalker; {
		if i%2 != 0 {
			// 说明是奇数节点

		} else {
			evenWalker.Next = walkerNode
			prev.Next = walkerNode.Next
			evenWalker = evenWalker.Next
		}
		i++
		prev = walkerNode
		walkerNode = walkerNode.Next

	}
	walkerNode := head
	for ; nil != walkerNode.Next; walkerNode = walkerNode.Next {

	}
	walkerNode.Next = evenDummy.Next
	return head
}
