/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/11 3:35 下午
# @File : lt_14_最长公共前缀.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 其实类似于暴力法,直接遍历全部,然后进行一个一个匹配即可
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ret := strs[0]
	for i := 1; i < len(strs); i++ {
		ret = longestCommonPrefixCompare(ret, strs[i])
	}
	return ret
}
func longestCommonPrefixCompare(str1, str2 string) string {
	l1, l2 := len(str1), len(str2)
	limit := l1
	if l2 < l1 {
		limit = l2
	}
	i := 0
	for ; i < limit; i++ {
		if str1[i] == str2[i] {
			continue
		}
		break
	}
	return str1[:i]
}
