/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-13 09:13 
# @File : lt_190_Reverse_Bits.go
# @Description : 反转二进制数
# @Attention : 思路就是讲数字最后一个数字一个一个取出,然后移动到首位
*/
package LeetCode

func reverseBits(num uint32) uint32 {
	index := 31
	sum := uint32(0)
	for num != 0 {
		lastOne := num & 1
		sum |= lastOne << uint32(index)
		num >>= 1
		index--
	}
	return sum
}
