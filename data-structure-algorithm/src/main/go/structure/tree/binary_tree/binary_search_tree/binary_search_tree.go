package binary_search_tree

import (
	"azure-container-networking/npm/util"
	"data-structure-algorithm/src/main/go/common"
	"fmt"
	"gopkg.in/eapache/queue.v1"
	"strconv"
)

/*
	二叉查找树的创建
	定义: 二叉查找树是指左孩子的值小于根节点,而右孩子的节点大于根节点的值
*/
type BinarySearchTreeNode struct {
	data       int // 为了方便演示,这里直接指定为int数据类型
	leftChild  *BinarySearchTreeNode
	rightChild *BinarySearchTreeNode
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
	size int
}

func (n *BinarySearchTreeNode) GetHashKey() int {
	hashString := util.Hash(fmt.Sprintf("%v", n.data))
	hashCode, e := strconv.Atoi(hashString)
	if nil != e {
		panic(e)
	}
	return hashCode
}
func NewBinarySearhTree() *BinarySearchTree {
	binarySearchTree := &BinarySearchTree{
		root: nil,
		size: 0,
	}
	return binarySearchTree
}

// 插入节点逻辑很简单,就是找到插入节点的父节点
// O(logn)遍历整颗树,如果小于父节点则遍历左子树,否则右子树,只需要判断是否为空即可
func (t *BinarySearchTree) InsertNode(value int) error {
	//if nil == value {
	//	return common.NilPointerError
	//}
	if nil == t.root {
		t.root = &BinarySearchTreeNode{data: value}
		return nil
	}
	newNode := &BinarySearchTreeNode{data: value}
	tempNode := t.root
	for {
		if newNode.data < tempNode.data {
			if nil == tempNode.leftChild {
				tempNode.leftChild = newNode
				break
			}
			tempNode = tempNode.leftChild
		} else {
			if nil == tempNode.rightChild {
				tempNode.rightChild = newNode
				break
			}
			tempNode = tempNode.rightChild
		}
	}
	t.size++
	return nil
}

// 删除节点的逻辑:
// 先遍历树找到这个节点,重点在于用什么去弥补这个节点的值,
//		1.若删除的节点没有右孩子,则用左孩子来代替作为根节点
// 		2.若被删除的节点的右孩子没有左孩子则直接用这个右孩子代替根节点
//		3.若被删除的节点的右孩子有左孩子,则用这个右孩子的左孩子作为根节点
//		总之: 先后顺序的判断:判断右孩子是否存在->存在判断右孩子的左孩子是否存在
//	具体的逻辑,因为涉及左右孩子,并且需要重新连接,所以需要获得value节点的父节点,因而我们需要一个节点保存上个节点的地址
//	先提前进行遍历,遍历结束之后对后续的值进行遍历
func (t *BinarySearchTree) DeleteNode(value int) error {
	tempNode := t.root
	var lastVisitNode *BinarySearchTreeNode // 保留的是父节点的地址
	for nil != tempNode && tempNode.data != value {
		lastVisitNode = tempNode
		if value < tempNode.data {
			tempNode = tempNode.leftChild
		} else {
			tempNode = tempNode.rightChild
		}
	}
	// 可能的情况: 到了末尾,从而tempNode为空了
	// 或者是匹配到了这个值
	if nil == tempNode {
		return common.NoSuchElementError
	}
	var newRootNode *BinarySearchTreeNode // 减少代码块,这里提前定义一个变量
	// 说明有这个值,此时lastVisitNode保存的是其父节点的值
	// 1. 判断右孩子是否存在
	if nil == tempNode.rightChild {
		// 不存在则直接将左孩子作为根节点
		// 此时是不需要对原先根节点的其他节点进行连接的,因为此时的根节点只有左孩子
		newRootNode = tempNode.leftChild
		// 2. 判断右孩子的的左孩子是否为空
	} else if nil == tempNode.rightChild.leftChild {
		// 如果右孩子的左孩子为空,则直接将这个右孩子作为根节点
		// 同时原先根节点的左孩子可能不为空,因而我们需要重新连接,但能确保的是原先节点的左孩子节点必定是比这个新节点小的
		// 因而直接赋值即可
		newRootNode = tempNode.rightChild
		newRootNode.leftChild = tempNode.leftChild
	} else {
		// 说明右孩子的左孩子不为空,则将这个右孩子的左孩子作为新的根节点
		newRootNode = tempNode.rightChild.leftChild
		// 此时他需要连接的元素:原先根节点的左孩子+原先根节点的右孩子
		// 而这个元素的原先位置也需要进行变更,因此此时需要递归进行处理
		t.DeleteNode(newRootNode.data)
		newRootNode.leftChild, newRootNode.rightChild = tempNode.leftChild, tempNode.rightChild
	}
	if value < lastVisitNode.data {
		lastVisitNode.leftChild = newRootNode
	} else {
		lastVisitNode.rightChild = newRootNode
	}
	return nil
}

func (t *BinarySearchTree) BFSTree() []int {
	results := make([]int, 0)
	queue := queue.New()

	if nil == t.root {
		return nil
	}
	queue.Add(t.root)
	for tempValue := queue.Remove(); ; tempValue = queue.Remove() {
		tempNode := tempValue.(*BinarySearchTreeNode)
		results = append(results, tempNode.data)
		if nil != tempNode.leftChild {
			queue.Add(tempNode.leftChild)
		}
		if nil != tempNode.rightChild {
			queue.Add(tempNode.rightChild)
		}
		if queue.Length() <= 0 {
			break
		}
	}
	return results
}
