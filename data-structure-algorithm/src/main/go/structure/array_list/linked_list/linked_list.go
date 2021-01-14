/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-12 10:03 
# @File : linked_list.go
# @Description : 
# @Attention : 
*/
package linked_list

import (
	"errors"
	"fmt"
)

type listNode struct {
	data interface{}
	next *listNode
}

func NewListNode(data interface{}) *listNode {
	return &listNode{data: data}
}

type linkedList struct {
	size int
	head *listNode
}

func NewLinkedList() *linkedList {
	l := new(linkedList)
	return l
}

func (this *linkedList) Add(data interface{}) {
	newNode := NewListNode(data)
	if this.head == nil {
		this.head = newNode
	} else {
		pHead := this.head
		if pHead.next != nil {
			pHead = pHead.next
		}
		pHead.next = newNode
	}
	this.size++
}

func (this *linkedList) RemoveByIndex(index int) (interface{}, error) {
	if index >= this.size {
		return nil, errors.New("index out of bound")
	}
	var data interface{}
	tempNode := this.head
	prev := this.head
	if index == 0 {
		data = this.head.data
		this.head = this.head.next
		return data, nil
	}
	for i := 0; i < index-1; i++ {
		prev = tempNode
		tempNode = tempNode.next
	}
	prev.next = tempNode.next
	return data, nil
}

func (this *linkedList) Show() {
	fmt.Println(this.head)
	iterator := this.Iterator()
	for {
		i, b := iterator()
		if !b {
			break
		}
		fmt.Println(i)
	}
}

func (this *linkedList) Iterator() func() (interface{}, bool) {
	tempNode := this.head
	return func() (interface{}, bool) {

		if tempNode == nil {
			return nil, false
		}
		data := tempNode.data
		tempNode = tempNode.next
		return data, true
	}
}

func (this *linkedList) Size() int {
	panic("implement me")
}
