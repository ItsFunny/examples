/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/4 9:33 上午
# @File : matrix_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	fmt.Println(updateMatrix([][]int{
		{0,0,0},
		{0,1,0},
		{0,0,0},
	}))
}
