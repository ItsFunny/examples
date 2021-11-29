/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/14 8:48 上午
# @File : lt_91_编码方法.go
# @Description :
# @Attention :
*/
package offer

// 状态转移方程:
// f(i)=f(i-1) + f(i-2)
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		// 代表着当取两位的时候,如果第一位为0 ,是无效的,然后得确保和<26才行
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+s[i-1]-'0' <= 26) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
