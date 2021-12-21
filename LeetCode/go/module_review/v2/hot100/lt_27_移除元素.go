/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/17 9:25 上午
# @File : lt_27_移除元素.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 就是 重新赋值
// func removeElement(nums []int, val int) int {
// 	index := 0
// 	for _, v := range nums {
// 		if v != val {
// 			nums[index] = v
// 			index++
// 		}
// 	}
// 	return index
// }

// 关键: 双指针,将相等的数,都放到后面去
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
