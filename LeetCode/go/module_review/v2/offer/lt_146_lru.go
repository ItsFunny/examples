/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/25 9:07 上午
# @File : lt_146_lru.go
# @Description :
# @Attention :
*/
package offer

type node struct {
	prev  *node
	next  *node
	value int
}

type DoubleLinkedList struct {
	head *node
	tail *node
}

type LRUCache struct {
	// key value
	dataM map[int]*node
	list  *DoubleLinkedList
	limit int
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		dataM: make(map[int]*node),
	}
	l := &DoubleLinkedList{
		head: &node{value: 0},
		tail: &node{value: 0},
	}
	l.head.prev = l.tail
	l.tail.next = l.head

	return ret
}

func (this *LRUCache) Get(key int) int {
	cache, exist := this.dataM[key]
	if !exist {
		return -1
	}
	this.list.moveToHead(cache)
	return cache.value
}
func (this *DoubleLinkedList) moveToHead(value *node) {
	removeNode(value)
	this.addToHead(value)
}
func removeNode(value *node) {
	next := value.next
	prev := value.prev
	prev.next = next
	next.prev = prev
	value.next, value.prev = nil, nil
}

func (this *DoubleLinkedList) addToHead(value *node) {
	next := this.head.next
	prev := this.head.prev

	value.next = next
	value.prev = prev

	next.prev = value
	prev.next = value

	this.head = value
}
func (this *DoubleLinkedList) removeTail() *node {
	prev := this.tail.prev
	next := this.tail.next
	next.prev = prev
	prev.next = next
	ret := this.tail
	this.tail = prev.next
	return ret
}
func (this *LRUCache) Put(key int, value int) {
	newNode := &node{
		prev:  nil,
		next:  nil,
		value: value,
	}
	origin, exist := this.dataM[key]
	if exist {
		this.list.moveToHead(origin)
	} else {
		// 计算长度
		this.list.addToHead(newNode)
		this.dataM[key] = newNode
		if len(this.dataM) > this.limit {
			this.list.removeTail()
		}
	}
}
