/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/2 9:18 上午
# @File : lt_279_完全平方数.go
# @Description :
# @Attention :
*/
package v2

import "math"

/*
给定正整数 n，
找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
完全平方数 是一个整数，
其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
*/

// 关键:
// 动态规划: 状态转移方程: f(n)=1+f(i-j*j)
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt32
		// 注意: 这里是 <= i ,因为我们需要统计的是 [0,n]的值
		for j := 1; j*j <= i; j++ {
			min = numSquaresMin(dp[i-j*j], min)
		}
		dp[i] = min + 1
	}
	return dp[n]
}

func numSquaresMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
