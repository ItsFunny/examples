/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/23 9:02 下午
# @File : lt_200_岛屿数量.go
# @Description :
# @Attention :
*/
package v2

// 关键:
// dfs
// 找到一个1 之后(既这是一个island), 此时要将其周围所有的1 都置为0 ,因为1连接成片的时候,是只算一个岛屿的
func numIslands2(grid [][]byte) int {
	ret := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				ret++
				numIslands2Dfs(grid, i, j)
			}
		}
	}

	return ret
}
func numIslands2Dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
	}
	numIslands2Dfs(grid, i-1, j)
	numIslands2Dfs(grid, i+1, j)
	numIslands2Dfs(grid, i, j-1)
	numIslands2Dfs(grid, i, j+1)
}
