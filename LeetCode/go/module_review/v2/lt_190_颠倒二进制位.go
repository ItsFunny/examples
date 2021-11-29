/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/22 10:26 上午
# @File : lt_190_颠倒二进制位.go
# @Description :
# @Attention :
*/
package v2

// 关键: 计算 1的个数
func reverseBits2(num uint32) uint32 {
	var ret uint32
	bitOneCount := 31
	for bitOneCount >= 0 {
		// x&1 判断bitOneCount 这个位置是否为1,然后左移,这样的话,如果最后一个为1 ,则可以得到二进制下的部分值
		ret += (num & 1) << bitOneCount
		bitOneCount--
		// 然后将num 消除最后一位置空 ,就可以不断的统计得到1的个数了(以及对应的位置)
		num >>= 1
	}
	return ret
}
