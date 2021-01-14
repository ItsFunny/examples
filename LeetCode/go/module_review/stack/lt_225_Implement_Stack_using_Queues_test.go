/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-11 09:18
# @File : lt_225_Implement_Stack_using_Queues.go
# @Description : 栈实现队列
# @Attention :
*/
package stack

import (
	"fmt"
	"testing"
)

func TestMyStack_Top(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
}
