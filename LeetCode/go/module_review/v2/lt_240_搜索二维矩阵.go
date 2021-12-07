/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/1 9:13 上午
# @File : lt_240_搜索二维矩阵.go
# @Description :
# @Attention :
*/
package v2

// 关键,从左下或者右上开始查询
func searchMatrix2(matrix [][]int, target int) bool {
	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
