/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/13 8:59 上午
# @File : jz_30_连续子数组的最大和.go
# @Description :
# @Attention :
*/
package offer

// 关键: dp[i]=max(arr[i],dp[i-1]+arr[i))
func FindGreatestSumOfSubArray(array []int) int {
	dp := make([]int, len(array))
	dp[0] = array[0]
	r := dp[0]
	for i := 1; i < len(array); i++ {
		dp[i] = FindGreatestSumOfSubArrayMax(array[i], dp[i-1]+array[i])
		r = FindGreatestSumOfSubArrayMax(dp[i], r)
	}
	return r
}
func FindGreatestSumOfSubArrayMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
