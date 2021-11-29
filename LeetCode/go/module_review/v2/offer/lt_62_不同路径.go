/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/3 8:47 上午
# @File : lt_62_不同路径.go
# @Description :
# @Attention :
*/
package offer

// 题目关键: 只能向右走,或者只能向下走
// 解题关键: (m,n)要么是从(m-1)过来,要么是从(n-1)过来,所以 f(m,n)=f(m-1,n)+f(m,n-1)
// 但是: 注意; 在第0行和第0列,只能向右走,或者是向下走,所以只有1种方法
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	// 向下走
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
