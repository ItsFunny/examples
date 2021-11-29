/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/29 9:24 上午
# @File : lt_238_除自身数组以外的乘积.go
# @Description :
# @Attention :
*/
package v2

// 关键: 当前数 = 左边的乘积 * 右边的乘积
// 如: 2,4,6,8  对于下标 2的值:  左边的乘积=2*4 右边的乘积=8  => 2*4*8
// 但是还有很关键的一点是,左边的起始为1 ,右边末尾结尾也是为1
func productExceptSelf(nums []int) []int {
	l, r, ret := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	// 左边的乘积,起始为0
	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = nums[i-1] * l[i-1]
	}
	// 右边的乘积,末尾为0
	r[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}

	// 最后计算最终值
	for i := 0; i < len(nums); i++ {
		// 当前结果的值为 左边的值* 右边的值,但是不要计算当前的值
		ret[i] = l[i] * r[i]
	}

	return ret
}
