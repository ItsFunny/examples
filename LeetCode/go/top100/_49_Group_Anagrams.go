/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-08 13:16 
# @File : _49_Group_Anagrams.go
# @Description : 字谜,将相同单词字母的放一起
# @Attention :
*/
package main

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, v := range strs {
		ana := getAna(v)
		if _, ok := dict[ana]; !ok {
			dict[ana] = make([]string, 0)
		}
		dict[ana] = append(dict[ana], v)
	}

	res := make([][]string, 0)
	for _, v := range dict  {
		res = append(res, v)
	}
	return res
}

func getAna(str string) string {
	bytes := make([]int, 128)
	for _, b := range str {
		bytes[b]++
	}
	builder := strings.Builder{}
	for i := 0; i < 128; i++ {
		if bytes[i] > 0 {
			builder.WriteByte(byte(i))
			builder.WriteString(string(bytes[i]))
		}
	}
	return builder.String()
}