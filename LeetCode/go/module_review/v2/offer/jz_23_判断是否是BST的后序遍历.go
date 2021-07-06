/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/6 9:20 上午
# @File : jz_23_判断是否是BST的后序遍历.go
# @Description :
# @Attention :
*/
package offer

// func VerifySquenceOfBST(sequence []int) bool {
// 	// write code here
// }
// func smallVerifySquenceOfBST(rootIndex int, sequence []int) bool {
// 	if rootIndex == 0 {
// 		return true
// 	}
// 	newRootIndex := -1
// 	for i := 0; i < rootIndex; i++ {
// 		if sequence[i] > sequence[rootIndex] {
// 			return false
// 		}
// 		if sequence[i] < sequence[rootIndex] {
// 			newRootIndex = i
// 		}
// 	}
//
// 	return true
// }
// func verifySquenceOfBST(rootIndex int, sequence []int) bool {
// 	var lastSmallIndex int
// 	for i := 0; i < rootIndex; i++ {
// 		if sequence[i] > sequence[rootIndex] {
// 			return false
// 		}
// 	}
// 	for i := rootIndex; i < len(sequence); i++ {
// 		if sequence[i] < sequence[rootIndex] {
// 			return false
// 		}
// 	}
// 	for index, v := range sequence {
// 		if v <= sequence[rootIndex] {
// 			lastSmallIndex = index
// 			rootIndex = index
// 			break
// 		}
// 	}
// }
