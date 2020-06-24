package binary_search_tree

import (
	"github.com/influxdata/influxdb/pkg/testing/assert"
	"testing"
)

/*
	The tree shall be like this
				     10
			5				   10			    删除10节点
		4      6		  		 18			------------->
				 8			15  	19
				/		  /	  \	      \
		 	7  			14     16		20
*/
var buildSearchTreeArray = []int{10, 5, 4, 6, 8, 7, 10, 18, 15, 19, 14, 16, 20}
var expcetedSuccessArray = []int{10, 5, 10, 4, 6, 18, 8, 15, 19, 7, 14, 16, 20}

// 测试删除的数据
var (
	deleteNodeTestArrayCase1 = []int{90, 50, 150, 20, 5, 25}
	case1ExceptedArray=[]int{90,20,150,5,25}
	case1DeleteValue=50
)
var(
	deleteNodeTestArrayCase2 = []int{90, 50, 150, 20, 125, 175, 5, 25, 140}
	case2ExceptedArray=[]int{90,50,175,20,125,5,25,140}
	case2DeleteValue=150
)
var (
	deleteNodeTestArrayCase3 = []int{90, 50, 150, 20, 75, 5, 66, 80, 68}
	case3ExceptedArray=[]int{90,66,150,20,75,5,68,80}
	case3DeleteValue=50
)

func TestBinarySearchTree_InsertNode(t *testing.T) {
	binarySearhTree := NewBinarySearhTree()
	for _, value := range buildSearchTreeArray {
		binarySearhTree.InsertNode(value)
	}
	assert.Equal(t, binarySearhTree.BFSTree(), expcetedSuccessArray)
}
// 1.若删除的节点没有右孩子,则用左孩子来代替作为根节点
// 删除的元素没有右节点
func TestBinarySearchTree_DeleteNodeCase1(t *testing.T) {
	binarySearhTree := NewBinarySearhTree()
	for _, value := range deleteNodeTestArrayCase1 {
		binarySearhTree.InsertNode(value)
	}
	binarySearhTree.DeleteNode(case1DeleteValue)
	assert.Equal(t,binarySearhTree.BFSTree(),case1ExceptedArray)
}
// 2.若被删除的节点的右孩子没有左孩子则直接用这个右孩子代替根节点
// 删除的元素有右孩子,但是右孩子的左孩子为空
func TestBinarySearchTree_DeleteNodeCase2(t *testing.T) {
	binarySearchTree:=NewBinarySearhTree()
	for _,value:=range deleteNodeTestArrayCase2{
		binarySearchTree.InsertNode(value)
	}
	binarySearchTree.DeleteNode(case2DeleteValue)
	assert.Equal(t,binarySearchTree.BFSTree(),case2ExceptedArray)
}

// 3.若被删除的节点的右孩子有左孩子,则用这个右孩子的左孩子作为根节点
// 删除的元素的右孩子不为空,且右孩子的左孩子也不为空(则左孩子的右孩子作为根节点)
func TestBinarySearchTree_DeleteNodeCase3(t *testing.T) {
	binarySearchTree:=NewBinarySearhTree()
	for _,value:=range deleteNodeTestArrayCase3{
		binarySearchTree.InsertNode(value)
	}
	binarySearchTree.DeleteNode(case3DeleteValue)
	assert.Equal(t,binarySearchTree.BFSTree(),case3ExceptedArray)
}