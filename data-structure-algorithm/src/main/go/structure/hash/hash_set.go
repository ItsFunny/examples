package hash

import (
	"fmt"
	"strconv"

	"azure-container-networking/npm/util"
)

// linkedList中的node节点
type LinkNode struct {
	data interface{}
	next *LinkNode
}

type LinkedList struct {
	head *LinkNode
	tail *LinkNode
	size int
}

type HashMap struct {
	buckets []*LinkedList
}

type HashSet struct {
	object interface{}
	setMap *HashMap
}

func (l *LinkedList)Push(data interface{}){
	newNode:=&LinkNode{
		data: data,
	}
	l.tail.next=newNode
	l.tail=newNode
}


func (m *HashMap)Put(key interface{},value interface{})bool{
	if nil==key {
		return false
	}
	hashCode, _ := strconv.Atoi(util.Hash(fmt.Sprintf("%v", key)))
	index:=hashCode% len(m.buckets)
	rootBucket:=m.buckets[index]

	if nil==rootBucket {
		newNode:=&LinkNode{
			data: value,
		}
		m.buckets[index]=&LinkedList{
			head: newNode,
			tail: newNode,
			size: 0,
		}
	}else{
		rootBucket.Push(value)
	}
	return true
}



func (s *HashSet)Add(key interface{})bool{
	return s.setMap.Put(key,s.object)
}

func NewHashSet()*HashSet{
	return &HashSet{
		object: struct{}{},
		setMap: &HashMap{
			buckets: make([]*LinkedList,16),
		},
	}
}

