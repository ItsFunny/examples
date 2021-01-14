/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-23 09:37 
# @File : lt_1438_Longest_Continuous_Subarray_With_Absolute_Diff_Less_Than_or_Equal_to_Limit.go
# @Description : 
# @Attention : 
*/
package slide_window

/*
	最大子串
	1. 遇到的问题, 以为是求最大子串,才知道是求子串中,最大和最小值<=limit 的时候的子串的最长子串
 */

func longestSubarray(nums []int, limit int) int {
	if len(nums) == 0 {
		return 0
	}
	// 确保 最后一位是当前的 left 和right 的下标
	maxQ := make([]int, 0)
	minQ := make([]int, 0)
	left, right := 0, 0
	result := 0
	for ; right < len(nums); right++ {
		for ; len(minQ) > 0 && minQ[nums[len(minQ)-1]] < nums[right]; right++ {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)
		for ; len(maxQ) > 0 && maxQ[nums[len(maxQ)-1]] > nums[right]; right++ {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)
		for len(maxQ) > 0 && len(minQ) > 0 && maxQ[len(maxQ)-1]-minQ[len(minQ)-1] > limit {
			left++
		}

		result = longestSubarrayMax(result, right-left+1)
	}

	return result
}

func longestSubarrayMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
