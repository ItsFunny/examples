/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/26 9:10 上午
# @File : lt_118_杨辉三角.go
# @Description :
# @Attention :
*/
package offer

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	ret := make([][]int, numRows)
	for index := range ret {
		ret[index] = make([]int, index+1)
		ret[index][0] = 1
		ret[index][index] = 1
		for j := 1; j < index; j++ {
			ret[index][j] = ret[index-1][j] + ret[index-1][j-1]
		}
	}
	return ret
}
