/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/16 9:44 下午
# @File : lt_7_整数反转.go
# @Description :
# @Attention :
*/
package offer

import "math"

func reverse(x int) int {
	// 123 => 321
	r := 0
	for x != 0 {
		if r > math.MaxInt32/10 || r < math.MinInt32/10 {
			return 0
		}
		// 先取得最后一个数
		last := x % 10 // 123 % 10 =3  => 12%10=2 => 1%10=1
		// 然后移除最后一个数
		x /= 10 // 123 /10 =12  12 /10 =1  => 1
		// 再累加结果
		r = r*10 + last // 0*10 +3  => 3*10+2=32  => 32*10 + 1=> 321
	}

	return r
}
