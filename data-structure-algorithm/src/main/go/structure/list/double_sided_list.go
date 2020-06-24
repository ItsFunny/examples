package list

import "data-structure-algorithm/src/main/go/common"

// 双向链表
// 指的是一个节点中既有指向前驱节点的指针,又有指向后继节点的指针

type DSListNode struct {
	data interface{}
	previous *DSListNode
	next *DSListNode
}

type DoubleSiededList struct {
	head *DSListNode
	tail *DSListNode
	size int
}

func (d *DoubleSiededList)Add(value interface{}){
	if nil==value{
		return
	}
	newNode:=&DSListNode{
		data:     value,
		}
	if d.size==0{
		d.head,d.tail=newNode,newNode
		d.tail.next=d.head
		d.head.previous=d.tail
	}else{
		d.tail.next=newNode
		newNode.previous=d.tail
		d.tail=newNode
		d.tail.next=d.head
	}
	d.size++
}

// @Function: 移除某个下标的元素
// 下标从0开始
func (d *DoubleSiededList)Remove(index int)error{
	if index>d.size {
		return common.IndexOutOfRangeError
	}
	// 判断是否是链头元素
	if index==0{
		d.head=d.head.next
		d.head.previous=d.tail
		d.tail.next=d.head
		d.size--
		return nil
	}else if index ==d.size-1{
		// 需要判断是否是链尾元素
		// 如果是链尾元素,则需要重新调整链尾
		d.tail=d.tail.previous
		d.tail.next=d.head
		d.head.previous=d.tail	// 别忘记头指向的前驱也需要更改
		d.size--
		return nil
	}
	// 删除指定下标的元素
	tempNode:=d.head
	for i:=0;i<index;i++{
		tempNode=tempNode.next
	}
	prevNode:=tempNode.previous
	prevNode.next=tempNode.next
	if nil!=tempNode.next{
		tempNode.next.previous=prevNode
	}
	d.size--
	return nil
}

// 删除指定value的节点
func (d *DoubleSiededList)DeleteConcreteValue(value interface{})error{
	if nil==value {
		return common.NilPointerError
	}
	if d.size==0 {
		return common.EmptyListError
	}
	if d.head.data==value {
		d.head.next.previous=d.tail
		d.head=d.head.next
		d.tail.next=d.head
	}
	tempNode:=d.head
	for {
		// 上面已经对头节点校验过了,直接跳到下一个即可
		tempNode=tempNode.next
		if tempNode.data==value {
			tempNode.previous.next=tempNode.next
			tempNode.next.previous=tempNode.previous
		}else if tempNode==d.head{
			return common.NoSuchElementError
		}
	}
	d.size--
	return nil
}

func (d *DoubleSiededList)CollectResults()[]interface{}{
	if d.size==0 {
		return nil
	}
	results:=make([]interface{},0)
	tempNode:=d.head
	results=append(results, tempNode.data)
	for {
		tempNode=tempNode.next
		if tempNode==d.head {
			return results
		}else{
			results=append(results,tempNode.data)
		}
	}
	return results
}

func NewDoubleSidedList()*DoubleSiededList{
	return &DoubleSiededList{
		head: nil,
		tail: nil,
		size: 0,
	}
}
