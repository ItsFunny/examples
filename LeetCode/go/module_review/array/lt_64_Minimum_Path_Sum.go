/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-30 09:15 
# @File : lt_64_Minimum_Path_Sum.go
# @Description : 
# @Attention : 
*/
package array

/*
	最小路径和
	依旧为动态规划题
 */

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			}
			if j == 0 && i > 0 {
				grid[i][j] += grid[i-1][j]
			}

			if i > 0 && j > 0 {
				// 取最小值
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
