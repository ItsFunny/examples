/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/19 9:00 上午
# @File : lt_14_最长公共前缀.go
# @Description :
# @Attention :
*/
package offer

// 最长公共前缀
// 关键: LCP(S1....SN)=LCP(LCP(LCP(S1,S2),S3)...SN)
// 既两两比较,得出最长之后,在继续比较即可
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	r := strs[0]
	for i := 1; i < len(strs); i++ {
		r = longestCommonPrefixWith(r, strs[i])
		if len(r) == 0 {
			break
		}
	}

	return r
}
func longestCommonPrefixWith(prefix, str2 string) string {
	l := len(prefix)
	if len(str2) < l {
		l = len(str2)
	}
	index := 0
	for ; index < l && prefix[index] == str2[index]; index++ {
	}
	return prefix[:index]
}
