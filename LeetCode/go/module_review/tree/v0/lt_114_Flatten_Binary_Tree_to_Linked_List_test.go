/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-09 08:14
# @File : lt_114_Flatten_Binary_Tree_to_Linked_List.go
# @Description :
# @Attention :
*/
package v0

import (
	"fmt"
	"testing"
)

func Test_flatten(t *testing.T) {
	r:=&TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
			},
			Right: &TreeNode{
				Val:   4,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{
				Val:   6,
			},
		},
	}
	r2:=&TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
			},
		},
	}
	flatten(r2)
	fmt.Println(r)
}
