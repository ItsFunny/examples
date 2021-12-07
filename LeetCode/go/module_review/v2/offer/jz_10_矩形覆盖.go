/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/30 9:06 上午
# @File : jz_10_矩形覆盖.go
# @Description :
# @Attention :
*/
package offer
// 动态规划,与 fiberN类似,f(n)=f(n-1)+f(n-2)
func rectCover(number int) int {
	if number <= 2 {
		return number
	}
	dp := make([]int, number)
	dp[0], dp[1] = 1, 2
	for i := 2; i < number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number-1]
}
