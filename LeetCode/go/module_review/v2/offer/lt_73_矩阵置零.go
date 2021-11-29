/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/6 8:37 下午
# @File : lt_73_矩阵置零.go
# @Description :
# @Attention :
*/
package offer

// 题目关键:
// 有一个元素为0,则所在行和列都为0
// 解题关键: 使用标记数组的方式,就是先遍历匹配,某个值为0,则标记该行为0
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for col, v := range matrix {
		for row, vv := range v {
			if vv == 0 {
				rows[col] = true
				cols[row] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if cols[j] || rows[i] {
				matrix[i][j] = 0
			}
		}
	}
}
