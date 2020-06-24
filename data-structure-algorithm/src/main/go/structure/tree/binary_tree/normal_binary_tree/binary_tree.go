package normal_binary_tree

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type TreeNode struct {
	data int
	leftChild *TreeNode
	rightChild *TreeNode
}
type BinaryTree struct {
	root *TreeNode
	index int
}
func (t *BinaryTree)LoopBuildTree(node *TreeNode,arr []int)*TreeNode{
	if t.index> len(arr)-1{
		return nil
	}
	if arr[t.index]==-1{
		t.index++
		return nil
	}else {
		node=&TreeNode{data:arr[t.index]}
		t.index++
		node.leftChild=t.LoopBuildTree(node.leftChild,arr)
		node.rightChild=t.LoopBuildTree(node.rightChild,arr)
		return node
	}
}

// 非递归创建普通二叉树,因为是给定的数组,所以以-1代表空,流程逻辑如下
// 核心就是二叉树节点满的时候只有2个节点
//  判断data==-1   true?说明要么插入右节点,要么是左右节点都没有,如果是插入右节点,则新node入队之后还需要修改方向为左
//				   false:判断插入的是左节点还是右节点,如果插入的是右节点,则需要修改方向为左(因为二叉树只有2个节点,
// 						  并且是先序创建:根左右) ,最后元素入队(因为可能后面的值不是-1,是要继续插入的)
//	ps: 在流程中并不需要对stack进行判空,因为开始之前先将根节点入队了,已经确保了不会为空
//  ps: 同时,其实二叉树是当遇到连续2个-1的时候会pop,那么n叉树则可以认为是遇到n个-1时是会出队的
//  ps : 因此,当n叉树时,判断的情况为: if counts!=n {} else{}
//  总结: 出队的情况: n叉树遇到n个连续的-1(子节点出栈,跳回父节点) | n叉树的节点满了(父节点出栈,子节点开始接收值)
func (t *BinaryTree)BuildTreeWithStack(arr []int){
	if len(arr)==0{
		return
	}
	t.root=&TreeNode{data:arr[0]}
	left:=true
	stack := arraystack.New()
	stack.Push(t.root)
	for i:=1;i< len(arr);i++{
		if arr[i]==-1{
			if left{
				// 重新调整方向
				left=false
			}else if !stack.Empty() {
				stack.Pop()
			}
		}else{
			node:=&TreeNode{data:arr[i]}
				temp, _ := stack.Peek()
				if left{
					temp.(*TreeNode).leftChild=node
				}else{
					temp.(*TreeNode).rightChild=node
					left=true
					stack.Pop()
				}
				stack.Push(node)
		}
	}
}


// iterator

// preOrder
func (t *BinaryTree)preOrderTree(node *TreeNode){
	if nil!=node{
		fmt.Printf("%d->",node.data)
		t.preOrderTree(node.leftChild)
		t.preOrderTree(node.rightChild)
	}
}

// 中序遍历  左根右
func (t *BinaryTree)InOrderIteratorTree(node *TreeNode){
	if nil!=node{
		t.InOrderIteratorTree(node.leftChild)
		fmt.Printf("%d->",node.data)
		t.InOrderIteratorTree(node.rightChild)
	}
}

// 后序遍历 左右根
// 这个node是根节点,所以按照左右根的顺序,先将左节点入队,再入右节点,最后打印根节点
func (t *BinaryTree)PostOrderInteratorTree(node *TreeNode){
	if nil!=node{
		t.PostOrderInteratorTree(node.leftChild)
		t.preOrderTree(node.rightChild)
		fmt.Printf("%d->",node.data)
	}
}

// BFS 遍历树,通过queue来实现
// 优点在于性能较DFS快,但是内存较DFS更大
func (t *BinaryTree)BFSTree(){

}