/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-31 15:31 
# @File : lt_95_Unique_Binary_Search_Trees_II.go
# @Description : 
# @Attention : 
*/
package v0

func generateTrees(n int) []*TreeNode {
	if n ==0 {
		return nil
	}
	return helper(1, n)
}

func helper(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var result []*TreeNode
	for i := start; i <= end; i++ {
		lefts := helper(start, i-1)
		rights := helper(i+1, end)
		for _, left := range lefts {
			for _, right := range rights {
				root := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				result = append(result, root)
			}
		}
	}
	return result
}
