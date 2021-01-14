/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-17 14:55 
# @File : lt_209_Minimum_Size_Subarray_Sum.go
# @Description : 
# @Attention : 
*/
package two_points

import "math"

/*
	找到  大于 s的最小组合
	如 7,并且提供了[2,3,1,2,4,3] ,则 4+3 >=7  2+3+2>=7 但是 3>2 所以 4,3为组合
	滑动窗口(双指针解决)
 */
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r, min := 0, 0, math.MaxInt32
	sum := 0
	for r < len(nums) {
		sum += nums[r]
		for sum >= s {
			//  开始缩小窗口
			min = minIn11(r - l + 1, min)
			sum-=nums[l]
			l++
		}
		r++
	}
	if min==math.MaxInt32{
		return 0
	}
	return min
}
func minIn11(n, m int) int {
	if n < m {
		return n
	}
	return m
}
