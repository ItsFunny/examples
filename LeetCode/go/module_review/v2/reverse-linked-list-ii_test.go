/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/26 10:01 上午
# @File : reverse-linked-list-ii_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_Rever2(t *testing.T) {
	node := CreateListNode(5)
	between := reverseBetween(node, 1, 3)
	fmt.Println(between.String())
}
