/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/22 10:12 上午
# @File : lt_189_轮转数组.go
# @Description :
# @Attention :
*/
package v2

// 关键:
// 3次翻转
// 第一次: 翻转全部
// 第二次: 翻转 0-k
// 第三次: 翻转k-剩下的
func rotate(nums []int, k int) {
	k %= len(nums)
	rotateReverse(nums)
	rotateReverse(nums[:k])
	rotateReverse(nums[k:])
}
func rotateReverse(nums []int) {
	limit := len(nums) >> 1
	l := len(nums) - 1
	for i := 0; i < limit; i++ {
		nums[i], nums[l-i] = nums[l-i], nums[i]
	}
}
