/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-30 09:34 
# @File : lt_73_Set_Matrix_Zeroes.go
# @Description : 
# @Attention : 
*/
package array

/*
	有0出现的行和列全变为0
	区分是手动设置为0的还是原先就为0的,所以需要一个 index 做运算
 */

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	// flag := 0
	modified := -100000
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == modified {
				continue
			}
			//
			// index := i*len(matrix[0]) + j + 1
			// if flag&index >= index || matrix[i][j] != 0 {
			// 	// 说明是手动赋值的
			// 	continue
			// }
			// 赋值为0 将 这一行的赋值为0
			for k := 0; k < len(matrix[0]); k++ {
				if k <= j {
					matrix[i][k] = 0
				} else if matrix[i][k] != 0 {
					// flag |= i*len(matrix[0]) + k + 1
					matrix[i][k] = modified
				}
			}
			// 将这一列赋值为0
			for k := 0; k < len(matrix); k++ {
				if k <= j {
					matrix[k][j] = 0
				} else if matrix[k][j] != 0 {
					// flag |= k*len(matrix[0]) + j + 1
					matrix[i][k] = modified
				}
			}
		}
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[i][j] == modified {
					matrix[i][j] = 0
				}
			}
		}
	}
}