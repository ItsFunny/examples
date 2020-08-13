/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-12 10:06 
# @File : lt_260_Single_Number_III.go
# @Description : 
# @Attention : 
*/
package byte

func singleNumber3(nums []int) []int {
	diff := 0

	// 将只出现1次的数据保留下来
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}

	// result中保存的是 只出现1次的数字,这个时候的数字是 a,b的混合数字,但是最后1位 1 要么是a的,要么是b的
	// 得到了最后一个数字
	result := []int{diff,diff}
	diff = (diff & (diff - 1)) ^ diff
	for i := 0; i < len(nums); i++ {
		if nums[i]&diff == 0 {
			// 这里再做一步 异或的原因在于 ,说明这个元素是a,则需要消除 a的公共部分,剩下的就是b了
			// 因为可能不仅a &diff==0 还有可能是出现2次的元素
			// 最终剩下的元素就是b了
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}
