/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-13 09:01 
# @File : lt_191_Number_of_1_Bits.go
# @Description : 
# @Attention : 
*/
package byte

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		count += int(num >> uint32(i) & uint32(1))
	}
	return count
}
