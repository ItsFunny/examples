/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-30 09:15
# @File : lt_64_Minimum_Path_Sum.go
# @Description :
# @Attention :
*/
package array

import (
	"fmt"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	fmt.Println(minPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}))
}
