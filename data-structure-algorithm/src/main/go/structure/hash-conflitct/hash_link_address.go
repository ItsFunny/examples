package hash_conflitct

import (
	"azure-container-networking/npm/util"
	"fmt"
	"strconv"
)

/*
	解决hash冲突的第二种方法:链地址法
 */

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type LinkedList struct {
	head *LinkNode
	tail *LinkNode
	size int
}

func (l *LinkedList) AddNode(value interface{}) {
	newNode := &LinkNode{data: value}
	if nil == l.head {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
	l.size++
}

func (l *LinkedList) PopNode() interface{} {
	if nil == l.head {
		return nil
	}
	value := l.head.data
	l.head = l.head.next
	return value
}

type HashMap struct {
	elements []*LinkedList
	size int
	capability int
	loadFactor float32
	threshold int
}

func (m *HashMap)Put(key interface{},value interface{}){
	if nil==key {
		return
	}
	hashString := util.Hash(fmt.Sprintf("%v", key))
	hashCode, _ :=strconv.Atoi(hashString)
	index:=hashCode&(len(m.elements)-1)
	rootNode:=m.elements[index]
	newNode:=&LinkNode{
		data: value,
	}
	if nil==rootNode{
		m.elements[index]=&LinkedList{
			head: newNode,
			tail: newNode,
			size: 1,
		}
	}else{
		rootNode.tail.next=newNode
		rootNode.tail=newNode
	}
	m.size++
	if m.size> m.threshold{
		// resize
	}
}

func (m *HashMap)resize(){
	//oldCapability:=m.capability
	//oldEles:=m.elements
	//oldThreshold:=m.threshold
	//newThreshold:=0


	// 重新建立阈值
	//if oldCapability >0 &&(oldThreshold<<1)<2^32-1{
	//	newThreshold=oldCapability<<1
	//}
	// 重新计算hash,然后复制,如果从头到尾遍历计算hash的话太慢了
}


func NewHashMap(maxEles int,loadFactor float32)*HashMap{
	hashMap:=&HashMap{
		elements:   make([]*LinkedList,0),
		size:       0,
		loadFactor: loadFactor,
	}
	return hashMap
}

