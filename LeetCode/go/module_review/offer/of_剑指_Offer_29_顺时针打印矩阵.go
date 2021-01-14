/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 15:57 
# @File : of_剑指_Offer_29_顺时针打印矩阵.go
# @Description : 
# @Attention : 
*/
package offer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 00 {
		return nil
	}

	rows := len(matrix)
	cols := len(matrix[0])
	result := make([]int, 0)
	left, right, top, bottom := 0, cols-1, 0, rows-1
	for {
		// left-> right
		for i := left; i <=right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// top -> bottom
		for i := top; i <=bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// right -> left
		for i := right; i >=left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if bottom<top{
			break
		}
		// bottom->top
		for i := bottom; i >=top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return result
}
