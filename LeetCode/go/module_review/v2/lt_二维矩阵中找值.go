/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/22 9:13 上午
# @File : lt_二维矩阵中找值.go
# @Description :
# @Attention :
*/
package v2

// 从左下或者右上入手
func searchMatrix(matrix [][]int, target int) bool {
	for i, j := len(matrix)-1, 0; i >= 0 && j < len(matrix[i]); {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}

	return false
}
