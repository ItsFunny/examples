/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/21 9:20 上午
# @File : lt_28_实现strStr().go
# @Description :
# @Attention :
*/
package hot100

// 只会暴力匹配,kmp什么吊玩意,对我没用,学个j
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	m, n := len(haystack), len(needle)
match:
	for i := 0; i+n < m; i++ {
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] {
				continue match
			}
			return i
		}
	}
	return -1
}
