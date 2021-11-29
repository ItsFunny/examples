/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/22 10:44 上午
# @File : lt_191_位1的个数.go
# @Description :
# @Attention :
*/
package v2

func hammingWeight2(num uint32) int {
	ret := 0
	for num > 0 {
		num = num & (num - 1)
		ret++
	}
	return ret
}
