/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/17 9:27 上午
# @File : lt_二进制_反转.go
# @Description :
# @Attention :
*/
package v2

func reverseBits(num uint32) uint32 {
	// 核心:
	// 将1不断的前移
	var r uint32
	var count = 31
	for num > 0 {
		r += (num & 1) << count
		count--
		num >>= 1
	}

	return r
}
