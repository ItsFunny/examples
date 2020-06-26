/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-15 11:00 
# @File : _17_Letter_Combinations_of_a_Phone_Number.go
# @Description : 手机键盘功能,排列组合问题,FIFO | MAP 能解决
# @Attention : 排列组合问题,都可以是直接用for循环遍历匹配
*/
package main

func letterCombinations(digits string) []string {
	num2letters := map[rune][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{}
	}
	opts := []string{""}
	for _, d := range digits {
		opts = updateOptions(opts, num2letters[d])
	}
	return opts
}

func updateOptions(opts []string, letters []string) []string {
	x := []string{}
	for _, o := range opts{
		for _, l := range letters{
			x = append(x, o+l)
		}
	}
	return x
}