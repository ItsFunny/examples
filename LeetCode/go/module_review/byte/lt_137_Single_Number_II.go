/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-12 09:38 
# @File : lt_137_Single_Number_II.go
# @Description : 
# @Attention : 
*/
package byte

func singleNumber2(nums []int) int {
	result := 0

	for i := 0; i < 64; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			// nums[j]>>i 的作用在于  判断nums[j]上的 第i个数是不是1
			count += nums[j] >> uint(i) & 1
		}
		result |= (count % 3) << uint(i)
	}
	return result
}
