/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/2 9:43 上午
# @File : number_of_islands.go
# @Description :
# @Attention :
*/
package v2

func numIslands(grid [][]byte) int {
	r := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && dfs(grid, i, j) >= 1 {
				r++
			}
		}
	}
	return r
}

func dfs(grid [][]byte, i int, j int) int {
	// 注意边界条件
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = 0
		return 1 + dfs(grid, i, j-1) + dfs(grid, i, j+1) + dfs(grid, i-1, j) + dfs(grid, i+1, j)
	}
	return 0
}
