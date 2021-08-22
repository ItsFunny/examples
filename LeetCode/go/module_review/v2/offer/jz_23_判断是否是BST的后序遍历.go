/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/6 9:20 上午
# @File : jz_23_判断是否是BST的后序遍历.go
# @Description :
# @Attention :
*/
package offer

func VerifySquenceOfBST(sequence []int) bool {
	if len(sequence)==0{
		return false
	}
	return verifySquenceOfBST(sequence, 0, len(sequence)-1)
}

func verifySquenceOfBST(sequence []int, start, end int) bool {
	if start >= end {
		return true
	}
	leftRootIndex := start
	for ; leftRootIndex <= end && sequence[leftRootIndex] < sequence[end]; leftRootIndex++ {}
	for i := leftRootIndex; i <= end; i++ {
		if sequence[i] < sequence[end] {
			return false
		}
	}
	return verifySquenceOfBST(sequence, start, leftRootIndex-1) && verifySquenceOfBST(sequence, leftRootIndex, end-1)
}
