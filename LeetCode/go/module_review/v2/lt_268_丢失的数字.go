/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/2 9:00 上午
# @File : lt_268_丢失的数字.go
# @Description :
# @Attention :
*/
package v2

// 关键: 数学公式: n*(n+1)/2
// 或者是排序,或者是hash表
func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	for _, v := range nums {
		total -= v
	}
	return total
}
