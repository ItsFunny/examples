/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-09 08:43 
# @File : lt_74_Search_a_2D_Matrix.go
# @Description : 
# @Attention : 
*/
package array

// 二维数组中查询一个数
/*
	从左到右递增
	从上到下递增
	从右上角开始
 */

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true;
		}
	}
	return false
}
