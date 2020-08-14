/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-13 09:28 
# @File : lt_201_Bitwise_AND_of_Numbers_Range.go
# @Description : 计算 [n,m]范围内的所有数字的按位与的和
# @Attention :
	思路呢,因为是按位与,所以一旦某个位置有0,必然为0
	因此只需要找到公共前缀即可
*/
package byte

func rangeBitwiseAnd(m int, n int) int {
	count := 0
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << uint(count)
}
