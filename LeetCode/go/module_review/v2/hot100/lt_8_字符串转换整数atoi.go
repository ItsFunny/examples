/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/10 6:31 下午
# @File : lt_8_字符串转换整数atoi.go
# @Description :
# @Attention :
*/
package hot100

import "math"

// 关键是多种极端情况下的测验
func myAtoi(s string) int {
	ret := 0

	// 1. 去除前导前缀
	index := 0
	for ; index < len(s); {
		if s[index] == ' ' {
			index++
			continue
		}
		break
	}
	// 2. 可能是极端情况,极端情况下会出现刚好达到长度
	if index == len(s) {
		return 0
	}
	// 判断是正数还是负数
	div := false
	if s[index] == '-' {
		div = true
		index++
	} else if s[index] == '+' {
		index++
	}
	for ; index < len(s); index++ {
		v := s[index]
		// 注意: 这一步是主要是为了,消除不合法的数,如 asdd aaa 123,这种是不合法的
		if v < '0' || v > '9' {
			break
		}
		if ret >= math.MaxInt32/10 {
			if div {
				return math.MinInt32
			}
			return math.MaxInt32
		} else if ret <= math.MinInt32/10 {
			if !div {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		ret = ret*10 + int(v-'0')
	}
	if div {
		ret *= -1
	}
	return ret
}
