/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-14 12:42 
# @File : list.go
# @Description : 
*/
package common

type Node struct {
	Data interface{} `json:"data"`
	Next *Node `json:"next"`
}

type LinkedList struct {
	Head *Node `json:"head"`
	Tail *Node `json:"tail"`
	Size int
}

func (list *LinkedList) Push(data interface{}) {
	newNode := &Node{Data: data}
	if nil == list.Head {
		list.Head = newNode
	} else {
		list.Tail.Next = newNode
	}
	list.Tail = newNode
	list.Size++
}

func (list *LinkedList) GetSize() int {
	return list.Size
}

func (list *LinkedList) Contains(value interface{}) bool {
	if list.Size == 0 {
		return false
	}

	tNode := list.Head
	for nil != tNode {
		if tNode.Data == value {
			return true
		}
		tNode = tNode.Next
	}

	return false
}
