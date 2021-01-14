/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-09 10:10 
# @File : _10_Regular_Expression_Matching.go
# @Description :   实现正则匹配
# @Attention : 
*/
package main


func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([][]bool, n + 1)
	for i := range(dp) {
		dp[i] = make([]bool, m + 1)
	}
	dp[0][0] = true
	if m > 0 && n > 0 {
		dp[1][1] = s[0] == p[0] || p[0] == '.'
	}
	for j:=2 ;j<=m; j++ {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	for i:=1; i<=n; i+=1 {
		for j:=2; j<=m; j+=1 {
			if p[j-1] != '*' {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			} else {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1]==p[j-2] || p[j-2] == '.'))
			}
		}
	}
	return dp[n][m]
}