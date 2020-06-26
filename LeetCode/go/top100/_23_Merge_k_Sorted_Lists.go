/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-18 10:33 
# @File : _23_Merge_k_Sorted_Lists.go
# @Description :
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
# @Attention :  实现排序队列,对其元素进行一个一个入队判断即可
*/
package main


type ByHead []*ListNode

func (self ByHead) Len() int { return len(self) }
func (self ByHead) Less(i,j int) bool {
	return self[i].Val < self[j].Val
}
func (self ByHead) Swap(i,j int) { self[i], self[j] = self[j], self[i] }

func (self *ByHead) Push(x interface{}) {
	*self = append(*self, x.(*ListNode))
}

func (self *ByHead) Pop() interface{} {
	h := *self
	el := h[len(h) - 1]
	*self = h[0:len(h) - 1]
	return el
}

func mergeKLists(lists []*ListNode) *ListNode {

	h := &ByHead{}

	for _,l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	if h.Len() == 0 {
		return nil
	}

	firstElement := heap.Pop(h).(*ListNode)

	var resHead *ListNode = firstElement
	var resTail *ListNode = firstElement

	if firstElement.Next != nil {
		heap.Push(h, firstElement.Next)
	}

	for h.Len() > 0 {

		currElement := heap.Pop(h).(*ListNode)

		resTail.Next = currElement
		resTail = resTail.Next

		if currElement.Next != nil {
			heap.Push(h, currElement.Next)
		}

	}

	return resHead


}