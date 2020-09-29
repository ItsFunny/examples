/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-29 09:16 
# @File : lt_62_Unique_Paths.go
# @Description : 
# @Attention : 
*/
package array

/*
	求路径的可能数
	第一思路: dfs,将访问过的标记为1 ,结果: dfs 会超时,只能使用动态规划

 */
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if 0 == i || 0 == j {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j - 1] + dp[i - 1][j]
			}
		}
	}
	return dp[m - 1][n - 1]
}
