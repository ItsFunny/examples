/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-17 09:50 
# @File : loop_linked_list.go
# @Description : 循环单链表
# @Attention : 
*/
package linked_list

import (
	"errors"
)

type CircleLinkedList struct {
	head *listNode
	size int
}

func (this *CircleLinkedList) Add(data interface{}) {
	newNode := NewListNode(data)
	this.size++
	if this.head == nil {
		this.head = newNode
	} else {
		pNode := this.head.next
		for pNode.next != this.head {
			pNode = pNode.next
		}
		pNode.next, newNode.next = newNode, this.head
	}
}

func (this *CircleLinkedList) RemoveByIndex(index int) (interface{}, error) {
	if index >= this.size {
		return nil, errors.New("arrayIndexOutOfBoundError")
	}
	pNode := this.head
	pPrev := this.head
	for i := 0; i < index-1; i++ {
		pPrev = pNode
		pNode = pNode.next
	}
	pPrev.next = pNode.next
	if pNode == this.head {
		this.head = pNode.next
	}
	return pNode.data, nil
}

func (this *CircleLinkedList) Show() func() (interface{}, bool) {
	pNode := this.head
	return func() (interface{}, bool) {
		if pNode == nil {
			return nil, false
		}
		data := pNode.data
		pNode = pNode.next
		return data, true
	}
}

func (this *CircleLinkedList) Size() int {
	return this.size
}

func NewCircleLinkedList() *CircleLinkedList {
	l := new(CircleLinkedList)
	return l
}
