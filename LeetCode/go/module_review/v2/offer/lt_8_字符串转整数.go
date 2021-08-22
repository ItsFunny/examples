/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/17 9:36 上午
# @File : lt_8_字符串转整数.go
# @Description :
# @Attention :
*/
package offer

import (
	"math"
)

func myAtoi(s string) int {
	r := 0
	div := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] == '-' {
			if div!=0{
				return 0
			}
			div = -1
			continue
		} else if s[i] == '+' {
			if div!=0{
				return 0
			}
			div = 1
			continue
		}
		v := s[i] - '0'
		if v < 0 || v > 9 {
			break
		}
		if r > math.MaxInt32/10 {
			if div == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		r = r*10 + int(v)
	}

	return r * div
}
