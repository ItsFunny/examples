/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-07 10:19 
# @File : _48_Rotate_Image.go
# @Description : 既横的变成竖的并且顺序反一下,
	1. 先 横的变成竖的,对角线交换
	2. 顺序反一下
# @Attention : 
*/
package main

func rotate(matrix [][]int) {
	// diagonal symmetry change
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// column symmetry change
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][i], matrix[j][len(matrix)-1-i] = matrix[j][len(matrix)-1-i], matrix[j][i]
		}
	}
}