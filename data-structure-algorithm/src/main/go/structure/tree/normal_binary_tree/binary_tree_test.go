package normal_binary_tree

import "testing"

// the tree should be :
/*
			1
		   / \
		 2    3
	   /   \  /  \
	4       5 6   7
 */
var arr=[]int{1,2,4,-1,-1,5,-1,-1,3,6,-1,-1,7}




// 递归创建测试,先序显示
func TestBinaryTree_LoopBuildTree(t *testing.T) {
	binaryTree:=&BinaryTree{}
	binaryTree.root=binaryTree.LoopBuildTree(nil,arr)
	binaryTree.preOrderTree(binaryTree.root)
}

// 非递归创建测试,先序显示
func TestBinaryTree_BuildTreeWithStack(t *testing.T) {
	binaryTree:=&BinaryTree{}
	binaryTree.BuildTreeWithStack(arr)
	binaryTree.preOrderTree(binaryTree.root)
}

// 中序遍历测试
func TestBinaryTree_InOrderIteratorTree(t *testing.T) {
	binaryTree:=&BinaryTree{}
	binaryTree.BuildTreeWithStack(arr)
	binaryTree.InOrderIteratorTree(binaryTree.root)
	// excepted to see :4,2,5,1,6,3,7
}

// 后序遍历测试
func TestBinaryTree_PostOrderInteratorTree(t *testing.T) {
	binartTree:=&BinaryTree{}
	binartTree.BuildTreeWithStack(arr)
	binartTree.PostOrderInteratorTree(binartTree.root)
	// excepted to see 4,5,2,6,7,3,1
}