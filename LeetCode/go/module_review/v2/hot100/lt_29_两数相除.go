/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/21 9:34 上午
# @File : lt_29_两数相除.go
# @Description :
# @Attention :
*/
package hot100

import "math"

// 关键: 除法可以等同为 加法,如 7/2 = 7比2大,至少为1,2翻倍为4,7比4大,则也翻倍,4翻倍为8比7大,所以要用7-4 与2继续比对
// 参考: https://leetcode-cn.com/problems/divide-two-integers/solution/po-su-de-xiang-fa-mei-you-wei-yun-suan-mei-you-yi-/
// 还要注意临界条件
func divide(dividend int, divisor int) int {
	flag := 1
	if dividend < 0 {
		flag *= -1
		dividend *= -1
	}
	if divisor < 0 {
		flag *= -1
		divisor *= -1
	}
	// 注意临界条件
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
func div(a, b int) int {
	if a < b {
		return 0
	}
	ret := 1
	tmp := b
	// 核心是这个for循环
	// 如 7/2 = 7比2大,至少为1,2翻倍为4,7比4大,则也翻倍,4翻倍为8比7大,所以要用7-4 与2继续比对
	for tmp+tmp <= a {
		ret = ret + ret
		tmp += tmp
	}
	return ret + div(a-tmp, b)
}
