package main

/*
 完全二叉树
 */


type CompleteTree struct {
	root *TreeNode
}

type TreeNode struct{
	data int
	leftChild *TreeNode
	rightChild *TreeNode
}

// 依据完全二叉树的特性
// 最下层的节点都在左边
// 之后[0,length/2]的有孩子节点
// 左孩子节点: 2*index ,右孩子节点下标: 2*index+1
func (t *CompleteTree)BuilcCompleteTreeWithStack(arr []int){
	trees:=make([]*TreeNode, len(arr))
	for _,data:=range arr{
		trees=append(trees,&TreeNode{data:data})
	}
	// 构建树
	for i:=0;i< len(arr)<<1;i++{
		if 2*i< len(arr){
			trees[i].leftChild=trees[2*i]
		}
		if 2*i+1< len(arr){
			trees[i].rightChild=trees[2*i+1]
		}
	}
}
