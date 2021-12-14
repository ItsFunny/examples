/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/13 8:35 下午
# @File : lt_17_电话号码的字母组合.go
# @Description :
# @Attention :
*/
package hot100

// 关键
// 1. 建立手机数字和字符的映射关系
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
	// 2 回溯法
	letterCombinationsBackTrack(digits, 0, "")
	return ret
}
func letterCombinationsBackTrack(digits string, index int, res string) {
	if index == len(digits) {
		ret = append(ret, res)
		return
	}
	str := phoneMap[digits[index]-'0']
	// 第三步: 各自遍历自己的数字所对应的值:回溯法, 如 2,3 ,当当前下标为0的时候,2对应的 value为: abc,3 对应的为:def
	// 则开始以 abc ,abd,abe 开始遍历
	for i := 0; i < len(str); i++ {
		// 最终得到的结果序列为: ad,ae,af,bd ...
		letterCombinationsBackTrack(digits, index+1, res+string(str[i]))
	}
}
