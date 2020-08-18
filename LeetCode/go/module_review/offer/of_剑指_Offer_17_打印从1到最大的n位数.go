/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 14:16 
# @File : of_剑指_Offer_17_打印从1到最大的n位数.go
# @Description : 
# @Attention : 
*/
package offer

import "strings"

var (
	sb     strings.Builder
	arrays []string
	max    int
	chars  = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

func printNumbers(n int) string {
	max = n
	return strings.Join(arrays, ",")
}

func dfs(index int) {
	if max == index {
		return
	}
	s := strings.Builder{}
	for i := 0; i < len(chars); i++ {
		s.WriteString(chars[i])
	}

}
