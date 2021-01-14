/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-21 08:53 
# @File : lt_54_Spiral_Matrix.go
# @Description : 
# @Attention : 
*/
package array

/*
	以蛇形的方式将二维数组打印出来
	从 左到右  => 从右到下=> 从下到左=> 从左到上
 */

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix)==0{
		return nil
	}
	if matrix[0]==nil{
		return nil
	}
	cols := len(matrix[0]) - 1
	rows := len(matrix) - 1
	left, top, right, bottom := 0, 0, cols, rows
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
