/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/11 10:04 上午
# @File : lt_求二叉树最大路径_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_aaa(t *testing.T){
	sum := maxPathSum(&TreeNode{
		Val: -3,
	})
	fmt.Println(sum)
}