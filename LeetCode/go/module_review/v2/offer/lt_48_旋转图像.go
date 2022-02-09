/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/24 10:09 下午
# @File : lt_48_旋转图像.go
# @Description :
# @Attention :
*/
package offer

// 关键:
// https://leetcode-cn.com/problems/rotate-image/solution/ji-qiao-ti-zai-zeng-song-yi-wei-xing-shi-377z/
// 1. 先对对角线进行翻转  两侧元素变化:  (0,1) => (1,0) ; (2,1)=>(1,2)
// 2. 再每一行中点进行翻转
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	// 以对角线进行翻转
	row := len(matrix)
	cow := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < cow; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 每一行的中点进行翻转
	mid := cow >> 1
	for i := 0; i < row; i++ {
		for j := 0; j < mid; j++ {
			matrix[i][j], matrix[i][cow-j-1] = matrix[i][cow-j-1], matrix[i][j]
		}
	}
}
