/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-14 10:01 
# @File : lt_74_search-a-2d-matrix.go
# @Description :  在二叉树矩阵中查询一个数
# @Attention : 将二维数组转为一维数组即可
注意,在获取中间值的时候,使用的是  matrix[mid/cols][mid%cols]  列的值
*/
package half

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	start := 0
	end := rows*cols - 1
	for start <= end {
		mid := start + (end-start)>>1
		midValue:=matrix[mid/cols][mid%cols]
		if midValue<target {
			start=mid+1
		}else if midValue>target{
			end=mid-1
		}else{
			return true
		}
	}
	return false
}
