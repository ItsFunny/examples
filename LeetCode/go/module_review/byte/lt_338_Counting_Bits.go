/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-13 09:07 
# @File : lt_338_Counting_Bits.go
# @Description :   就是统计 [0,num]数之间的 1出现的次数
# @Attention : 
*/
package byte

func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		result[i] = count(i)
	}

	return result
}

func count(num int) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
