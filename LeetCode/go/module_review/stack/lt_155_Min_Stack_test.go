/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-07 09:41
# @File : lt_155_Min_Stack_test.go.go
# @Description :
# @Attention :
*/
package stack

import (
	"fmt"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	stack.Top()
	stack.GetMin()
}
