package list

import "errors"

type ListNode struct {
	data interface{}
	next *ListNode
}

type SingleList struct {
	root *ListNode
}

func (l *SingleList) AddNode(value interface{}) {
	if nil == value {
		return
	}
	newNode := &ListNode{
		data: value,
	}
	if nil == l.root {
		l.root = newNode
		return
	}
	tmpNode := l.root
	for nil != tmpNode.next {
		tmpNode = tmpNode.next
	}
	tmpNode.next = newNode
}

// 删除尾部
func (l *SingleList) Remove() {
	tmpNode := l.root
	for nil != tmpNode.next && nil != tmpNode.next.next {
		tmpNode = tmpNode.next
	}
	tmpNode.next = nil
}

// 链表的删除需要之前删除节点的前驱节点
// 删除的时候要先连再断 (既先将前驱节点与这个节点的后继节点相连接)
func (l *SingleList) RemoveByValue(value interface{}) (bool, error) {
	tmpNode := l.root
	if tmpNode.data == value {
		l.root = tmpNode.next
		return true, nil
	}
	for tmpNode.next != nil && tmpNode.next.data != value {
		tmpNode = tmpNode.next
	}
	// tmpNode.next=nil | tmpNode.next.value=value
	if nil == tmpNode.next {
		return false, errors.New("元素不存在")
	}
	tmpNode.next = tmpNode.next.next
	return true, nil
}

func (l *SingleList) RemoveByIndex(index int) (bool, error) {
	if index <= 0 {
		return false, errors.New("no such index,根节点从1开始算起")
	} else if index == 1 {
		l.root = l.root.next
		return true, nil
	}
	tmpNode := l.root
	// 获取他的前驱节点
	for i := 1; i < index-1; i++ {
		if nil == tmpNode.next {
			// 不额外定义一个变量来判断
			return false, errors.New("index超过链表的长度")
		}
	}
	tmpNode.next = tmpNode.next.next
	return true, nil
}

func (l *SingleList) IteratorNode() []interface{} {
	values := make([]interface{}, 0)
	tmpNode := l.root
	for nil != tmpNode {
		values = append(values, tmpNode.data)
		tmpNode = tmpNode.next
	}
	return values
}
