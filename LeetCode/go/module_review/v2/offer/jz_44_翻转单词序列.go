/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/22 8:53 上午
# @File : jz_44_翻转单词序列.go
# @Description :
# @Attention :
*/
package offer

import "strings"

func ReverseSentence(str string) string {
	splits := strings.Split(str, " ")
	for i, j := 0, len(splits)-1; i < j; {
		splits[i], splits[j] = splits[j], splits[i]
		i++
		j--
	}
	return strings.Join(splits, " ")
}
