/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/31 9:28 上午
# @File : evaluate_reverse_polish_notation_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	evalRPN([]string{"4", "13", "5", "/", "+"})
}

func Test_Slice(t *testing.T) {
	bytes := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(bytes[:3])

	fmt.Println(bytes[1:6])
}
