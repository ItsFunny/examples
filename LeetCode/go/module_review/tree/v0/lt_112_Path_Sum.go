/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-08 08:14 
# @File : lt_112_Path_Sum.go
# @Description : 
# @Attention : 
*/
package v0

/*
	判断是否存在路径相加的和为该值
	依旧为递归
	注意点: 必须到叶子节点才可以
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if nil == root {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
