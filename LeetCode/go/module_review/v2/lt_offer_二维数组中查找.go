/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/25 9:32 上午
# @File : lt_offer_二维数组中查找.go
# @Description :
# @Attention :
*/
package v2

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix)==0{
		return false
	}
	for left, down := 0, len(matrix[0])-1; left < len(matrix) && down >= 0; {
		if matrix[left][down]>target{
			down--
		}else if matrix[left][down]<target{
			left++
		}else{
			return true
		}
	}
	return false
}
