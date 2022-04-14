package hot100

// 关键: 还是回溯算法
// 对于当前位置i,从 (0,i-1) 构建左子树, 从(i+1,n)构建右子树
// 最后从左子树中选择一棵,右子树中选择一棵
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesHelper(1, n)
}
func generateTreesHelper(start, end int) []*TreeNode {
	// 开始递归构建
	if start > end {
		// 注意这里, 当不符合条件的时候,不能直接返回nil,而是需要返回一个空的节点
		return []*TreeNode{nil}
	}
	ret := []*TreeNode{}
	for i := start; i <= end; i++ {
		// 所有可行的左子树
		leftTrees := generateTreesHelper(start, i-1)
		// 所有可行的右子树
		rightTrees := generateTreesHelper(i+1, end)

		// 然后从左和右中选择一棵
		for _, lt := range leftTrees {
			for _, rt := range rightTrees {
				cur := &TreeNode{Val: i}
				cur.Left = lt
				cur.Right = rt
				ret = append(ret, cur)
			}
		}
	}

	return ret
}
