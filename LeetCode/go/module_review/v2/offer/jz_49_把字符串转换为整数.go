/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/26 9:17 上午
# @File : jz_49_把字符串转换为整数.go
# @Description :
# @Attention :
*/
package offer

func StrToInt(str string) int {
	negative := false
	if len(str) == 0 {
		return 0
	}
	if str[0] == '-' {
		negative = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	r := 0
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return 0
		}
		r = r*10 + int(str[i]-'0')
	}
	if negative {
		r *= -1
	}
	return r
}
