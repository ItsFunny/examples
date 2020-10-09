/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-09 08:43
# @File : lt_74_Search_a_2D_Matrix.go
# @Description :
# @Attention :
*/
package array

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	matrix:=[][]int{
		{1,3,5,7},
		{10,11,16,20},
		{23,30,34,50},
	}
	b := searchMatrix(matrix, 3)
	fmt.Println(b)
}
