/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/16 9:40 上午
# @File : lt_二进制_出现1次的2个数.go
# @Description :
# @Attention :
*/
package v2

func singleNumber3(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	mix := 0
	for _, v := range nums {
		// 消除相同的数
		mix ^= v
	}
	r := []int{mix, mix}
	// 此时在mix 中的数,是 结果 a,b的混合数,并且最后一个1 ,基于异或的特性 a^b 最后一个1 要么是a的要么是b的
	// 假设是a的
	mix = (mix & (mix - 1)) ^ mix
	for _, v := range nums {
		if v&mix == 0 {
			r[0] ^= v
		} else {
			// 进入的条件: 可能是 a ,也可能是那些重复的元素 (重复的元素的话,基于异或的特性,会将重复的都消除)
			// 如果是a,则会将a 消除,所以r[1]最终会是 b
			r[1] ^= v
		}
	}
	return r
}
