/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 15:25 
# @File : of_剑指_Offer_26_树的子结构.go
# @Description : 判断是不是子树
# @Attention : 
*/
package offer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return nil != A && B != nil && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}
func recur(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return recur(a.Left, a.Left) && recur(a.Right, b.Left)
}
