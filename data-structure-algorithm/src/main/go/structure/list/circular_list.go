package list

import "errors"

// 循环链表

/*
	注意点:
		1. 删除的时候需要对是否是尾节点进行判断(如果是尾节点需要重新调整尾节点的指向)
		2. 设定一个size变量来充当长度
		3. 当遍历的时候,值是不可能为空的,因为是一个循环的链表,永远不会有空的指向,所以判断是否到头要通过判断是否与tail指针
相等
*/
type CircularNode struct {
	data interface{}
	next *CircularNode
}

type CircularList struct {
	head *CircularNode
	tail *CircularNode
}

func (l *CircularList) AddNode(value interface{}) {
	newNode := &CircularNode{data: value}
	if nil == l.head {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
	l.tail.next = l.head
}

// 尾部删除
func (l *CircularList) Remove() (interface{}, error) {
	var removeValue interface{}
	if nil == l.head {
		return nil, errors.New("the list is empty")
	}
	tempNode := l.head
	for tempNode.next != l.tail {
		tempNode = tempNode.next
	}
	removeValue = tempNode.next.data
	l.tail = tempNode
	l.tail.next = l.head
	return removeValue, nil
}

func (l *CircularList) RemoveByIndex(index int) (interface{}, error) {
	var removeValue interface{}
	var tempNode = l.head
	if nil == l.head {
		return nil, errors.New("NullPointerException")
	} else if tempNode == l.tail {
		// 表示只有1个元素
		removeValue = tempNode.data
		l.head, l.tail = nil, nil
		return removeValue, nil
	}
	if index <= 0 {
		return nil, errors.New("index invalid,root index is 1")
	} else if index == 1 {
		removeValue = tempNode.data
		l.head = tempNode.next
		l.tail.next = l.head
		return removeValue, nil
	} else {
		for i := 1; i < index; i++ {
			// 循环链表,这里不可能为nil,
			// 循环链表的删除需求与一般链表的删除操作不同,因为是一个环,所以会一直循环直到一个次数
			// 这里可以通过 index=index %size 进行优化,减少次数
			tempNode = tempNode.next
		}
	}
	// 判断这个值是否存在  xxxx 这个值是必然存在的
	//if nil == tempNode.next {
	//	return nil, errors.New("IndexOutOfBoundException")
	//}
	removeValue = tempNode.next.data
	// 判断下一个节点是否是尾节点
	if l.tail == tempNode.next {
		// 如果是尾节点,意味着需要移动尾节点的位置
		l.tail = tempNode
		l.tail.next = l.head
	} else {
		// 如果不是尾节点则直接移除
		tempNode.next = tempNode.next.next
	}
	return removeValue, nil
}

func (l *CircularList) IterateNode() ([]interface{}, error) {
	if nil == l.head {
		return nil, errors.New("NullPointerException")
	}
	tempNode := l.head
	values := make([]interface{}, 0)
	for nil != tempNode && tempNode != l.tail {
		values = append(values, tempNode.data)
		tempNode = tempNode.next
	}
	values = append(values, tempNode.data)
	return values, nil
}


