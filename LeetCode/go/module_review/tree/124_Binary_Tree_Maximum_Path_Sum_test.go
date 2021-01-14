/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-09 09:42
# @File : 124_Binary_Tree_Maximum_Path_Sum.go
# @Description :   求树的最大路径和
	分治法
	1. 左子树的最大路径
	2. 右子树的最大路径
	3. 左+右子树的最大路径
	4. 取最大值
# @Attention :
*/
package tree

import (
	"fmt"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(maxPathSum(root))
}
