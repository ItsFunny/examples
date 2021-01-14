/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-29 08:50 
# @File : lt_59_Spiral_Matrix_II.go
# @Description : 
# @Attention : 
*/
package array

/*
	给定一个数,生成二维数组
 */

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}
	max := n * n
	left, right, top, bottom := 0, n-1, 0, n-1

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	index := 1
	for index <= max {
		for i := left; i <= right; i++ {
			result[top][i] = index
			index++
		}
		top++
		for i := top; i <= bottom; i++ {
			result[i][right] = index
			index++
		}
		right--
		for i := right; i >=left; i-- {
			result[bottom][i] = index
			index++
		}
		bottom--
		for i := bottom; i >=top; i-- {
			result[i][left] = index
			index++
		}
		left++
	}
	return result
}
