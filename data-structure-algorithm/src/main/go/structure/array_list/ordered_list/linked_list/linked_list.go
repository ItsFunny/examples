/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-12 10:03 
# @File : linked_list.go
# @Description : 
# @Attention : 
*/
package linked_list

type listNode struct {
	data interface{}
	next *listNode
}

func NewListNode(data interface{}) *listNode {
	return &listNode{data: data}
}

type linkedList struct {
	size int
	data *listNode
}

func NewLinkedList() *linkedList {
	l := new(linkedList)
	return l
}

func (this *linkedList) Add(data interface{}) {
	newNode := NewListNode(data)
	if this.data == nil {
		this.data = newNode
	} else {
		pHead := this.data
		if pHead.next != nil {
			pHead = pHead.next
		}
		pHead.next = newNode
	}
	this.size++
}

func (this *linkedList) RemoveByIndex(index int) interface{} {
	panic("implement me")
}

func (this *linkedList) Show() func() (interface{}, bool) {
	panic("implement me")
}

func (this *linkedList) Size() int {
	panic("implement me")
}
