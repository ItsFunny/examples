/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/2 9:42 上午
# @File : lt_283_移动0.go
# @Description :
# @Attention :
*/
package v2

// 关键: 可以将非0的数,都放到前面,以修改索引的方式
func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index++
	}
}
