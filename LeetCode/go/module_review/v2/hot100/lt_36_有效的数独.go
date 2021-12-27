/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/27 9:00 上午
# @File : lt_36_有效的数独.go
# @Description :
# @Attention :
*/
package hot100

// 关键:
// 根据题意: 一行数字不可以相同,一列数字也不可以相同,斜线也不可以相同
// 一次遍历+ 子盒子的下标计算为 i/3,j/3
func isValidSudoku(board [][]byte) bool {
	var (
		// 代表的是,9行,每行的9个元素的次数
		rows [9][9]int
		// 代表的是9列,每列9个元素的次数
		cols [9][9]int

		// 在每个 3*3 的子box中,9个数出现的次数
		subboxes [3][3][9]int
	)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}
			// index 为 num-'0',因为数字从0开始,所以直接减去1 ,等价于 num-'0'-1
			index := num - '1'
			// 更新行中出现的次数
			rows[i][index]++
			// 更新列
			cols[j][index]++

			// 更新子盒子内的值
			subboxes[i/3][j/3][index]++

			// 最后计算结果,因为只能出现1次,所以>1 就是false
			if rows[i][index] > 1 || cols[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
