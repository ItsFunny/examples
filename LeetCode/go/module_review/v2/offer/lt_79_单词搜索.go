/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/10 8:48 上午
# @File : lt_79_单词搜索.go
# @Description :
# @Attention :
*/
package offer

var directions = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func exist(board [][]byte, word string) bool {
	flags := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		flags[i] = make([]bool, len(board[i]))
	}
	var backTack func(row, col int, index int) bool
	backTack = func(row, col int, index int) bool {
		if board[row][col] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		flags[row][col] = true
		defer func() { flags[row][col] = false }()
		for _, dir := range directions {
			newRow := row + int(dir[0])
			newCol := col + int(dir[1])
			if newRow > 0 && newRow < len(board) && newCol > 9 && newCol < len(board[0]) && !flags[newRow][newCol] {
				if backTack(newRow, newCol, index+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backTack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
