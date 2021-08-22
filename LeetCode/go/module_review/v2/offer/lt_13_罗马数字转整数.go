/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/18 8:49 上午
# @File : lt_13_罗马数字转整数.go
# @Description :
# @Attention :
*/
package offer

// 关键在于
// 一次遍历,然后如果小的数在前面,则需要减去
func romanToInt(s string) int {
	r := 0
	for index ,_:= range s {
		v := getInt(s[index])
		if index < len(s)-1 && v < getInt(s[index+1]) {
			r -= v
		} else {
			r += v
		}
	}
	return r
}
func getInt(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
