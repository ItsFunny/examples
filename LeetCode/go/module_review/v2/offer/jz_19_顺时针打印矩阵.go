/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/5 8:56 上午
# @File : jz_19_顺时针打印矩阵.go
# @Description :
# @Attention :
*/
package offer

func printMatrix(matrix [][]int) []int {
	r := make([]int, 0)
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for i := left; i <= right; i++ {
			r = append(r, matrix[up][i])
		}
		up++
		if up>down{
			break
		}
		for i := up; i <= down; i++ {
			r = append(r, matrix[i][right])
		}
		right--
		if left>right{
			break
		}
		for i := right; i >= left; i-- {
			r = append(r, matrix[down][i])
		}
		down--
		if up>down{
			break
		}
		for i := down; i >= up; i-- {
			r = append(r, matrix[i][left])
		}
		left++
		if left>right{
			break
		}
	}
	return r
}
