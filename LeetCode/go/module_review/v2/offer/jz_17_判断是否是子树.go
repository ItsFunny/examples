/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/4 1:37 下午
# @File : jz_17_判断是否是子树.go
# @Description :
# @Attention :
*/
package offer

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2 == nil || pRoot1 == nil {
		return false
	}
	// 因为可能这个节点是root的根节点,或者是root的left节点,也可能是root的right节点
	return isChild(pRoot1, pRoot2) || HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}
func isChild(root *TreeNode, child *TreeNode) bool {
	if root == nil {
		return false
	}
	if child == nil {
		return true
	}
	if root.Val != child.Val {
		return false
	}
	return isChild(root.Left, child.Left) && isChild(root.Right, child.Right)
}
