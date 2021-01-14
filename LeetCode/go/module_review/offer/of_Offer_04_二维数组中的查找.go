/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-14 10:18 
# @File : of_Offer_04_二维数组中的查找.go
# @Description :
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
# @Attention :
*/
package offer

func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	start := 0
	end := rows*cols - 1
	for start <= end {
		mid := start + (end-start)>>1
		midValue := matrix[mid/cols][mid%cols]
		if midValue < target {
			start = mid + 1
		} else if midValue > target {
			end = mid - 1
		} else {
			return true
		}
	}
	return false
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		} else {
			return true
		}
	}
	return false
}
