/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/26 8:48 上午
# @File : lt_36_有效的数独.go
# @Description :
# @Attention :
*/
package offer

// 参考: https://leetcode-cn.com/problems/valid-sudoku/solution/you-xiao-de-shu-du-by-leetcode/
// 关键:
// 1次遍历即可
// 格子的计算公式为: box_index := (i/3)*3 + j/3
func isValidSudoku(board [][]byte) bool {
	rows := make([][]byte, 9)
	cols := make([][]byte, 9)
	boxs := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		rows[i] = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0}
		cols[i] = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0}
		boxs[i] = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			value := board[i][j] - '1'
			if board[i][j] == '.' {
				continue
			}
			box_index := (i/3)*3 + j/3

			if rows[i][value] > 0 || cols[j][value] > 0 || boxs[box_index][value] > 0 {
				return false
			}
			rows[i][value]++
			cols[j][value]++
			boxs[box_index][value]++
		}
	}
	return true
}
