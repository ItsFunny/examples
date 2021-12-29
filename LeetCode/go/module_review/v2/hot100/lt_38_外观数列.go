/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/28 8:51 上午
# @File : lt_38_外观数列.go
# @Description :
# @Attention :
*/
package hot100

import (
	"strconv"
	"strings"
)

/*
n=5的时候:
1
11
21
1211
111221
一步一步来
给一个数，这个数是1
描述上一步的数，这个数是 1 即一个1，故写作11
描述上一步的数，这个数是11即两个1，故写作21
描述上一步的数，这个数是21即一个2一个1，故写作12-11
描述上一步的数，这个数是1211即一个1一个2两个1，故写作11-12-21
*/
// 关键是计算 n中出现相同数字的次数
func countAndSay(n int) string {
	// 初始为1
	prev := "1"
	count := 0
	for i := 2; i <= n; i++ {
		cur := strings.Builder{}
		// start=j 可以使得跳到下一个不匹配的开始,如之前是 11234,则start=j 可以跳到2开始
		for j, start := 0, 0; j < len(prev); start = j {
			//开始计算次数
			for j < len(prev) && prev[j] == prev[start] {
				j++
				count++
			}
			cur.WriteString(strconv.Itoa(count))
			cur.WriteByte(prev[start])
			count = 0
		}
		prev = cur.String()
	}
	return prev
}
