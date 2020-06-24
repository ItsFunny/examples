package list_algorithm

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

var nonLoopArray=[]int{1,2,3,4,5,6,7}
var loopArray=[]int{1,2,3,4,5,1,6,7}

func TestLinkedList_ValidIfLoop_False(t *testing.T) {
	linkedList:=&LinkedList{}
	for _,value:=range nonLoopArray{
		linkedList.Push(value)
	}
	assert.Equal(t,linkedList.ValidIfLoop(),false)
}

func TestLinkedList_ValidIfLoop_True(t *testing.T) {
	node1:=&ListNode{data:1}
	node2:=&ListNode{data:2}
	node3:=&ListNode{data:3}
	node4:=&ListNode{data:4}
	node5:=&ListNode{data:5}
	node1.next,node2.next,node3.next,node4.next,node5.next=node2,node3,node4,node5,node2
	linkedList:=&LinkedList{root:node1}
	linkedList.size=5

	assert.Equal(t,linkedList.ValidIfLoop(),true)
}

// 通过map判断是否有环
func TestLinkedList_ValidIfLoopByMap(t *testing.T) {
	node1:=&ListNode{data:1}
	node2:=&ListNode{data:2}
	node3:=&ListNode{data:3}
	node4:=&ListNode{data:4}
	node5:=&ListNode{data:5}
	node1.next,node2.next,node3.next,node4.next,node5.next=node2,node3,node4,node5,node2
	linkedList:=&LinkedList{root:node1}
	linkedList.size=5

	assert.Equal(t,linkedList.ValidIfLoopByMap(),true)
}

// 获取回环的节点
func TestLinkedList_GetLoopNode(t *testing.T) {
	node1:=&ListNode{data:1}
	node2:=&ListNode{data:2}
	node3:=&ListNode{data:3}
	node4:=&ListNode{data:4}
	node5:=&ListNode{data:5}
	node1.next,node2.next,node3.next,node4.next,node5.next=node2,node3,node4,node5,node2
	linkedList:=&LinkedList{root:node1}
	linkedList.size=5

	assert.Equal(t,linkedList.GetLoopNode(),node2)
}

// 只有2个节点的情况
func TestLinkedList_ValidIfLoop_TwoNodes(t *testing.T) {
	node1:=&ListNode{data:1}
	node2:=&ListNode{data:2}

	node1.next=node2
	node2.next=node1
	linkedList:=&LinkedList{root:node1}
	linkedList.size=2
	assert.Equal(t,linkedList.ValidIfLoop(),true)
}