/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/6 8:29 上午
# @File : lt_70_爬楼梯.go
# @Description :
# @Attention :
*/
package offer

// 解题关键: 动态规划方程式: f(n)=f(n-1)+f(n-2)
func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = 1
			continue
		}
		if i == 2 {
			dp[i] = 2
			continue
		}
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
