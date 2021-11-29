/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/1 8:29 下午
# @File : lt_54_螺旋矩阵.go
# @Description :
# @Attention :
*/
package offer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	// cols := len(matrix[0]) - 1
	// rows := len(matrix) - 1
	// left, top, right, bottom := 0, 0, cols, rows
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	result := make([]int, 0)
	for ; left <= right && top <= bottom; {
		// 因为起始是要获取到left 首位的,所以只需要获取首位即可
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		if left < right && top < bottom {
			for i := right - 1; i > left; i-- {
				result = append(result, matrix[bottom][i])
			}
			for i := bottom; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		top++
		right--
		bottom--
	}
	return result
}
