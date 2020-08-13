/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-12 09:10 
# @File : lt_136_Single_Number.go
# @Description :   找出那个为单次出现的数字
# @Attention : 
*/
package byte

func singleNumber(nums []int) int {

	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}

	return result
}
