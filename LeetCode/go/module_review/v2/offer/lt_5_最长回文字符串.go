/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/16 9:16 上午
# @File : lt_5_最长回文字符串.go
# @Description :
# @Attention :
*/
package offer

// 关键: 动态规划
// 如果去掉首尾,依旧为回文串,并且首尾相同
// s(i,j)=s(i+1,j-1) ^ (Si==Sj)
// 当len=1, true
// 当len=2 & s[0]==s[1] true
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return string(s[0])
	}
	dps := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dps[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			dps[i][j] = true
		}
	}

	// 最大长度
	maxL := 0
	// 最长回文串的截取开始位置
	subBegin := 0

	// 开始填值, 并且是一列一列填值
	for j := 1; j < len(s); j++ {
		for i := 0; i < len(s); i++ {
			if s[i] != s[j] {
				// 如果首尾不相同,肯定不是回文字符串
				dps[i][j] = false
			} else {
				l := j - i + 1
				if l < 3 {
					// 如果长度小于3 ,肯定是回文字符串
					dps[i][j] = true
				} else {
					// 长度大于3 ,则通过子串 来判断是否是回文字符串
					dps[i][j] = dps[i+1][j-1]
				}
			}
			// 更新结果集
			if dps[i][j] && j-i+1 > maxL {
				maxL = j - i + 1
				subBegin = i
			}
		}
	}

	return s[subBegin : subBegin+maxL]
}
