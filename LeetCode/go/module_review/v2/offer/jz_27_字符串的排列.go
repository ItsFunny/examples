/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/9 8:49 上午
# @File : jz_27_字符串的排列.go
# @Description :
# @Attention :
*/
package offer


func permutationDfs(strBytes []byte) []string {
	if len(strBytes) == 0 || len(strBytes) == 1 {
		return []string{string(strBytes)}
	}
	r := make([]string, 0)
	r = append(r)

	return r
}
