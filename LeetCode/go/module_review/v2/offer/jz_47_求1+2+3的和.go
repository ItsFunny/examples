/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/26 8:59 上午
# @File : jz_47_求1+2+3的和.go
# @Description :
# @Attention :
*/
package offer

func Sum_Solution(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
	}
	return dp[len(dp)-1]
}
