/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/22 9:08 下午
# @File : lt_198_打家劫舍.go
# @Description :
# @Attention :
*/
package v2

// 动态规划:
//  dp[0]=nums[0] , dp[n]=max(dp[n-2]+nums[n],dp[n-1])
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = robMax(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = robMax(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(dp)-1]
}

func robMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
