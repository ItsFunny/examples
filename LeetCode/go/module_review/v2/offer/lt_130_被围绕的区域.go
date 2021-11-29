/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/12 8:57 上午
# @File : lt_130_被围绕的取余.go
# @Description :
# @Attention :
*/
package offer

// 解题关键
// 边界上的点不会被填充为x ,则 从正方形4个边界,进行dfs
var (
	n, m int
)

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m = len(board), len(board[0])

	for i := 0; i < n; i++ {
		solveDfs(board, i, 0)
		solveDfs(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		solveDfs(board, 0, i)
		solveDfs(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}
func solveDfs(board [][]byte, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	solveDfs(board, x+1, y)
	solveDfs(board, x-1, y)
	solveDfs(board, x, y+1)
	solveDfs(board, x, y-1)
}
