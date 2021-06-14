/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-16 07:02 
# @File : lt_707_Design_Linked_List.go
# @Description : 
# @Attention : 
*/
package v0

/*
	实现linkedlist
	双向链表解决
 */
type MyLinkedList struct {
	first *LinkedNode
	last *LinkedNode
}

type LinkedNode struct {
	val int
	next *LinkedNode
	prev *LinkedNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{nil, nil}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	node, _ := this.getNode(index)
	if node == nil {
		return -1
	}
	return node.val
}

func (this *MyLinkedList) getNode(index int) (*LinkedNode, int) {
	if index < 0 {
		i := -1
		n := this.first
		for n != nil && i > index {
			n = n.prev
			i--
		}
		return n, i
	} else {
		i := 0
		n := this.first
		for n != nil && i < index {
			n = n.next
			i++
		}
		return n, i
	}
}

func (this *MyLinkedList) AddAtHead(val int)  {
	if this.first == nil {
		// add node to empty list
		this.first = &LinkedNode{val: val, next: nil, prev: nil}
		this.last = this.first
	} else {
		n := &LinkedNode{val: val, next: this.first, prev: nil}
		this.first.prev = n
		this.first = n
	}
}

func (this *MyLinkedList) AddAtTail(val int)  {
	if this.last == nil {
		this.AddAtHead(val)
	} else {
		n := &LinkedNode{val: val, next: nil, prev: this.last}
		this.last.next = n
		this.last = n
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index == 0 {
		this.AddAtHead(val)
	} else {
		node, i := this.getNode(index)
		if node == nil {
			if i == index {
				this.AddAtTail(val)
			}
		} else {
			// add new node before "node"
			n := &LinkedNode{val: val, next: node, prev: node.prev}
			node.prev.next = n
			node.prev = n
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 {
		// are you kidding me
		return
	}
	n, _ := this.getNode(index)
	if n == nil {
		return
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n == this.first {
		this.first = n.next
	}
	if n == this.last {
		this.last = n.prev
	}
}
/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
