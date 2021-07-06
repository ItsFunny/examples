/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/29 5:09 下午
# @File : linked_list_cycle_ii_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	node := CreateNodeBy(3, 2, 0, -4)
	cycle := detectCycle(node)
	fmt.Println(cycle)
}

func TestOther(t *testing.T) {
	fmt.Println(2 ^  2)
}
