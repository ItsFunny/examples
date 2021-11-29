/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/24 8:36 上午
# @File : lt_29_两数相除.go
# @Description :
# @Attention :
*/
package offer

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	flag := 1
	if dividend < 0 {
		flag *= -1
		dividend *= -1
	}
	if divisor < 0 {
		flag *= -1
		divisor *= -1
	}
	if divisor == 1 {
		ret := dividend * flag
		if ret > math.MaxInt32 {
			return math.MaxInt32
		} else if ret < math.MinInt32 {
			return math.MinInt32
		}
		return ret
	}

	return div(dividend, divisor) * flag
}
func div(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}
	r := 1
	tmpDivisor := divisor
	for tmpDivisor+tmpDivisor <= dividend {
		r = r + r
		tmpDivisor = tmpDivisor + tmpDivisor
	}
	return r + div(dividend-tmpDivisor, divisor)
}
