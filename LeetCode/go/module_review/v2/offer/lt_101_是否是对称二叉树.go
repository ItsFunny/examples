/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/18 9:03 上午
# @File : lt_101_是否是对称二叉树.go
# @Description :
# @Attention :
*/
package offer

// 解题思路:
/*
	判断是否是对称二叉树
	则判断子树是否是对称二叉树(递归)
	判断子树是否是对称二叉树=>左边的和右边的对称(想到双指针)
*/

func isSymmetric(root *TreeNode) bool {
	return checkIsSymmetric(root, root)
}
func checkIsSymmetric(r1, r2 *TreeNode) bool {
	if nil == r1 && nil == r2 {
		return true
	}
	if nil == r1 || nil == r2 {
		return false
	}
	return r1.Val == r2.Val && checkIsSymmetric(r1.Left, r2.Right) && checkIsSymmetric(r1.Right, r2.Left)
}
