/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-30 09:34
# @File : lt_73_Set_Matrix_Zeroes.go
# @Description :
# @Attention :
*/
package array

import "testing"

func Test_setZeroes(t *testing.T) {
	setZeroes([][]int{
		{0, 1, 2,0},
		{3, 4, 5,2},
		{1, 3, 1,5},
	})
}
