/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-11 09:07 
# @File : lt_number_of_islands.go
# @Description : 
# @Attention : 
*/
package stack

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}
	return count
}
func dfs(grid [][]byte, i, j int) int {
	if i < 0 || i >= len(grid) || j >= len(grid[i]) || j < 0 {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		return dfs(grid, i, j-1) + dfs(grid, i, j+1) + dfs(grid, i-1, j) + dfs(grid, i+1, j) + 1
	}
	return 0
}
