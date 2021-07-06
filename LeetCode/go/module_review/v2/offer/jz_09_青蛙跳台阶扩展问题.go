/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/30 8:51 上午
# @File : jz_09_青蛙跳台阶扩展问题.go
# @Description :
# @Attention :
*/
package offer

func jumpFloorII(number int) int {
	// write code here
	dp := make([]int, number)
	for i := 0; i < number; i++ {
		if i == 0 {
			dp[0] = 1
			continue
		}
		dp[i] = 2 * dp[i-1]
	}
	return dp[len(dp)-1]
}
