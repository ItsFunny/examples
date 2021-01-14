/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 09:07 
# @File : lt_394_Decode_String.go
# @Description : 
# @Attention : 
*/
package stack

import (
	"strconv"
	"strings"
)

// 给定一个经过编码的字符串，返回它解码后的字符串。
// s = "3[a]2[bc]", 返回 "aaabcbc". s = "3[a2[c]]", 返回 "accaccacc".
// s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	sb := strings.Builder{}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			vars := make([]byte, 0)
			for {
				temp := stack[len(stack)-1]
				if temp == '[' {
					stack = stack[:len(stack)-1]
					break
				}
				stack = stack[:len(stack)-1]
				vars = append(vars, temp)
			}
			// 反转
			reverse(&vars)
			// 获取弹出次数
			nums := make([]byte, 0)
			for len(stack) > 0 {
				temp := stack[len(stack)-1]
				if temp <= 57 && temp >= 48 {
					// 说明是数字
					nums = append(nums, temp)
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			sb.Reset()
			// 弹出次数
			reverse(&nums)
			sb.Write(nums)
			count, _ := strconv.Atoi(sb.String())
			for j := 0; j < count; j++ {
				stack = append(stack, vars...)
			}
		default:
			stack = append(stack, s[i])
		}
	}
	sb.Reset()
	sb.Write(stack)
	return sb.String()
}

func reverse(bytes *[]byte) {
	bs := *bytes
	for i, j := 0, len(bs)-1; i < j; {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}
