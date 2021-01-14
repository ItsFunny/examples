/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-08 08:30 
# @File : lt_113_Path_Sum_II.go
# @Description : 
# @Attention : 
*/
package tree

/*
	找到路径的同时,收集路径
	解题思路: dfs解决
 */

func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	dfs(root, sum, []int{}, &result)
	return result
}

func dfs(root *TreeNode, sum int, path []int, result *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		newPath := make([]int, len(path))
		copy(newPath, path)
		*result = append(*result, newPath)
		return
	}
	dfs(root.Left, sum-root.Val, path, result)
	dfs(root.Right, sum-root.Val, path, result)
}
