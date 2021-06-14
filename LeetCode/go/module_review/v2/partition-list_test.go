/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/28 10:00 上午
# @File : partition-list_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_Csss(t *testing.T) {
	by := CreateNodeBy(1, 4, 3, 2, 5, 2)
	fmt.Println(by.String())
	node := partition(by, 3)
	fmt.Println(node.String())
}
