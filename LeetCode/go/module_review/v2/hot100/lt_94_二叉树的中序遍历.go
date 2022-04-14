package hot100

// 关键: 中序遍历: 左根右
// 栈遍历或者是递归遍历
func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalWithStack(root)
}

// 栈遍历
func inorderTraversalWithStack(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || nil != root {
		for nil != root {
			// 栈的话,先要把所有的左节点都入栈
			stack = append(stack, root)
			root = root.Left
		}
		// 然后弹出
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		// 注意这一步,这时候要用弹出的node 的右节点去遍历
		root = node.Right
	}

	return ret
}

// 递归遍历
func inorderTraversalWithLoop(root *TreeNode) []int {
	ret := make([]int, 0)
	doInorderTraversalWithLoop(&ret, root)
	return ret
}
func doInorderTraversalWithLoop(ret *[]int, root *TreeNode) {
	if nil == root {
		return
	}
	if nil != root.Left {
		doInorderTraversalWithLoop(ret, root.Left)
	}
	*ret = append(*ret, root.Val)
	if nil != root.Right {
		doInorderTraversalWithLoop(ret, root.Right)
	}
}
