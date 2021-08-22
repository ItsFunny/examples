/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/20 9:04 上午
# @File : lt_17_电话号码组合.go
# @Description :
# @Attention :
*/
package offer

var phoneMap map[byte]string = map[byte]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}
var ret []string

// 关键: 回溯法
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	ret = nil
	letterCombinationsBackTrack(digits, 0, "")
	return ret
}
func letterCombinationsBackTrack(digits string, index int, res string) {
	if index == len(digits) {
		ret = append(ret, res)
		return
	}
	str := phoneMap[digits[index]-'0']
	for i := 0; i < len(str); i++ {
		letterCombinationsBackTrack(digits, index+1, res+string(str[i]))
	}
}
