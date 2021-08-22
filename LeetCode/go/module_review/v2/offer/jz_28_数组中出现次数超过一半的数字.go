/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/12 9:28 上午
# @File : jz_28_数组中出现次数超过一半的数字.go
# @Description :
# @Attention :
*/
package offer

func MoreThanHalfNum_Solution(numbers []int) int {
	// 关键: 排序
	numQuickSort(numbers, 0, len(numbers)-1)
	return numbers[len(numbers)>>1]
}

func numQuickSort(nums []int, left, right int) {
	if left < right {
		paration := numsGetParation(nums, left, right)
		numQuickSort(nums, left, paration)
		numQuickSort(nums, paration+1, right)
	}
}
func numsGetParation(nums []int, left, right int) int {
	standard := nums[left]
	for left < right {
		for ; right > left && nums[right] >= standard; {
			right--
		}
		nums[left] = nums[right]
		for ; left < right && nums[left] <= standard; {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left]=standard
	return left
}
