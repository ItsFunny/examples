/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-13 09:14
# @File : _98_Validate_Binary_Search_Tree.go
# @Description : 判断是否是二叉搜索树
	什么是二叉搜索树: 左孩子 < root < 右孩子
# @Attention :
*/
package v0

import "testing"

func Test_isValidBST(t *testing.T) {
	// type args struct {
	// 	root *TreeNode
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want bool
	// }{
	// 	struct {
	// 		name string
	// 		args args
	// 		want bool
	// 	}{
	// 		name: "",
	// 		args: args{
	// 			root: &TreeNode{
	// 				Val: 2,
	// 				Left: &TreeNode{
	// 					Val: 1,
	// 				},
	// 				Right: &TreeNode{
	// 					Val:   3,
	// 					Right: ndil,
	// 				},
	// 			},
	// 		},
	// 		want: true,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := isValidBST(tt.args.root); got != tt.want {
	// 			t.Errorf("isValidBST() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	isValidBST(root)
}
