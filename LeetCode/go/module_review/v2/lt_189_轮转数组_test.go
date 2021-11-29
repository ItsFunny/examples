/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/22 10:18 上午
# @File : lt_189_轮转数组_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_rotateReverse(t *testing.T) {
	arrs := []int{1, 2, 3, 4, 5, 6, 7}
	rotateReverse(arrs)
	fmt.Println(arrs)
}
