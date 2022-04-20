package hot100

// 关键:
// 树的特点:
// 左子树<root<右子树
// dfs: 对每棵子树判断是否是正确的二叉搜索树
// 左边最大值 要小于node ,右边最小值要>node
func isValidBST(root *TreeNode) bool {

	type tempResult struct {
		max *TreeNode
		min *TreeNode
		valid bool
	}
	var dfs func(node *TreeNode) tempResult

	dfs = func(node *TreeNode) tempResult {
		ret := tempResult{}
		ret.valid = true
		if node == nil {
			return ret
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if !left.valid || !right.valid {
			ret.valid=false
			return ret
		}
		// 左边最大值 要小于node ,右边最小值要>node
		if (left.max!=nil && left.max.Val>=node.Val) || (nil!=right.min && right.min.Val<=node.Val){
			ret.valid=false
			return ret
		}
		ret.max=node
		if right.max!=nil{
			ret.max=right.max
		}
		ret.min=node
		if left.min!=nil{
			ret.min=left.min
		}
		return ret
	}
	return dfs(root).valid
}
