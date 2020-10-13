/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-13 08:43 
# @File : lt_79_Word_Search.go
# @Description : 
# @Attention : 
*/
package array

/*
	dfs
	陷阱: 元素的组成必须是相邻的,所以不是什么滑动窗口题
 */

// func exist(board [][]byte, word string) bool {
// 	have := make([]int, 128)
// 	need := make([]int, 128)
// 	for i := 0; i < len(word); i++ {
// 		need[word[i]]++
// 	}
//
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			have[board[i][j]]++
// 		}
// 	}
// 	for i:=0;i<len(word);i++{
// 		if have[word[i]]<need[word[i]]{
// 			return false
// 		}
// 	}
// 	return true
// }
//
// var (
// 	direction = [][]int{
// 		{-1, 0},
// 		{1,0},
// 		{0,-1},
// 		{0, 1},
// 	}
// )
//
// func exist(board [][]byte, word string) bool {
// 	if len(board) == 0 || len(word) == 0 {
// 		return false
// 	}
// 	marker := make([][]bool, len(board))
// 	for i := 0; i < len(marker); i++ {
// 		marker[i] = make([]bool, len(board[i]))
// 	}
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[0]); j++ {
// 			if existDfs(board, i, j, 0, word, marker) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
//
// func existDfs(board [][]byte, i, j, wordIndex int, word string, marked [][]bool) bool {
// 	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
// 		return false
// 	}
// 	if wordIndex == len(word)-1 {
// 		return board[i][j] == word[wordIndex]
// 	}
// 	if board[i][j] == word[wordIndex] {
// 		marked[i][j] = true
// 		// 上下左右移动
// 		// 上
// 		for k:=0;k<4;k++{
// 			newI:=i+direction[k][0]
// 			newJ:=j+direction[k][1]
// 			if existDfs(board, newI, newJ, wordIndex+1, word, marked) && !marked[i][j] {
// 				return true
// 			}
// 		}
// 		marked[i][j]=false
// 	}
// 	return false
// }

var dirs = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if dfs(board, word, i, j) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] != word[0] {
		return false
	}

	temp := board[i][j]
	board[i][j] = '#'
	for _, dir := range dirs {
		ii := i + dir[0]
		jj := j + dir[1]

		if dfs(board, word[1:], ii, jj) {
			return true
		}
	}

	board[i][j] = temp
	return false
}
