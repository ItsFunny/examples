/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/23 9:13 上午
# @File : jz_45_扑克牌顺子.go
# @Description :
# @Attention :
*/
package offer

// 判断是否是连续的数字
func IsContinuous(numbers []int) bool {
	// write code here
	IsContinuousQSort(numbers, 0, len(numbers)-1)
	last := 0
	firstIndex:=-1
	for index, v := range numbers {
		if v == 0 {
			continue
		}
		if firstIndex==-1{
			firstIndex=index
		}
		if last == v {
			return false
		}
		last = v
	}
	return numbers[len(numbers)-1]-numbers[firstIndex] < 5
}
func IsContinuousQSort(nums []int, left, right int) {
	if left < right {
		paration := IsContinuousQSortParation(nums, left, right)
		IsContinuousQSort(nums, left, paration)
		IsContinuousQSort(nums, paration+1, right)
	}
}

func IsContinuousQSortParation(nums []int, left int, right int) int {
	standard := nums[left]
	for left < right {
		for ; right > left && nums[right] >= standard; right-- {
		}
		nums[left] = nums[right]
		for ; left < right && nums[left] <= standard; left++ {
		}
		nums[right] = nums[left]
	}
	nums[left] = standard
	return left
}
