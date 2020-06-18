/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-18 09:46 
# @File : doubly_linked_list.go
# @Description : 不带头双向链表
# @Attention : 
*/
package linked_list

type doubleListNode struct {
	prev *doubleListNode
	next *doubleListNode
	data interface{}
}

func NewDoublyListNode(data interface{}, prev *doubleListNode, next *doubleListNode) *doubleListNode {
	return &doubleListNode{
		prev: prev,
		next: next,
		data: data,
	}
}

type DoublyLinkedList struct {
	size int
	Head *doubleListNode
}

func (this *DoublyLinkedList) Add(data interface{}) {
	this.InsertHead(data)
}

func (this *DoublyLinkedList) RemoveByIndex(index int) (interface{}, error) {
	panic("implement me")
}

func (this *DoublyLinkedList) Show() func() (interface{}, bool) {
	panic("implement me")
}

func (this *DoublyLinkedList) Size() int {
	panic("implement me")
}

// 头插法
func (this *DoublyLinkedList) InsertHead(data interface{}) {
	newNode := NewDoublyListNode(data, nil, nil)
	if this.Head == nil {
		this.Head = newNode
		newNode.next = this.Head
	} else {
		newNode.next = this.Head
		this.Head.prev.next = newNode
		this.Head = newNode
	}
	this.size++
}

func (this *DoublyLinkedList) InsertTail(data interface{}) {
	panic("implement me")
}
