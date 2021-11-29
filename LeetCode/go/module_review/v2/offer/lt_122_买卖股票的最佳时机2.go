/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/29 9:09 上午
# @File : lt_122_买卖股票的最佳时机2.go
# @Description :
# @Attention :
*/
package offer

// 二维数组
// 动态规划:行代表的是第n天,列 0代表的是股票卖了后的收益,1代表的是买入股票的收益
func maxProfit2(prices []int) int {
	ret := 0
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	// 初始状态,代表的是第0天,买入股票的收益,此时为 负
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		// 此时卖出去股票,收益取最大值
		dp[i][0] = maxProfit2Max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 此时买入股票,更新今天的收益
		dp[i][1] = maxProfit2Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	ret = dp[len(prices)-1][0]
	return ret
}
func maxProfit2Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
