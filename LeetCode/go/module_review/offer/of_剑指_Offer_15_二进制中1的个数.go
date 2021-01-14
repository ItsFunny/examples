/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 14:11 
# @File : of_剑指_Offer_15_二进制中1的个数.go
# @Description : 
# @Attention : 
*/
package offer

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num & 1)
		num >>= 1
	}
	return count
}
