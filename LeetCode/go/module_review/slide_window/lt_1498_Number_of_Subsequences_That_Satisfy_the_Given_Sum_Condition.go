/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-27 09:00 
# @File : lt_1498_Number_of_Subsequences_That_Satisfy_the_Given_Sum_Condition.go
# @Description : 
# @Attention : 
*/
package slide_window

/*
	陷入一个误区, 以为是区间内的所有值
	原来只是最小值和最大值相加即可
	所以是排列组合问题即可
 */

func numSubseq(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	numSubseqQsort(nums, 0, len(nums)-1)
	if nums[0]*2 > target {
		return 0
	}

	left, right := uint(0), uint(len(nums)-1)
	res := 0
	for left <= right {
		if nums[left]+nums[right] <= target {
			res += 1 << (right - left)
			left++
		} else {
			right--
		}
	}

	return res % 1000000007
}

func numSubseqQsort(nums []int, start, end int) {
	if start < end {
		paration := numSubseqQsortParation(nums, start, end)
		numSubseqQsort(nums, start, paration)
		numSubseqQsort(nums, paration+1, end)
	}
}

func numSubseqQsortParation(nums []int, start int, end int) int {
	standard := nums[start]
	for start < end {
		for ; end > start && nums[end] >= standard; end-- {
		}
		nums[start] = nums[end]
		for ; start < end && nums[start] <= standard; start++ {
		}
		nums[end] = nums[start]
	}
	nums[start] = standard
	return start
}
