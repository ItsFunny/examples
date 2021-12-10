/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/8 9:38 下午
# @File : lt_7_整数反转.go
# @Description :
# @Attention :
*/
package hot100

import "math"

func reverse(x int) int {
	ret := 0
	div := false
	if x < 0 {
		x = x * -1
		div = true
	}
	for x > 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	if div {
		ret = ret * -1
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}
	return ret
}

