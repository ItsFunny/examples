/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-21 19:59
# @File : lt_56_Merge_Intervals.go
# @Description :
# @Attention :
*/
package array

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	a := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	a = [][]int{
		[]int{1, 4},
		[]int{1, 4},
	}
	merge(a)
	fmt.Println(a)
}
