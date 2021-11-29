/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/5 10:00 上午
# @File : lt_69_x的平方根.go
# @Description :
# @Attention :
*/
package offer

func mySqrt(x int) int {
	l, r := 0, x
	ret := 0
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ret
}
