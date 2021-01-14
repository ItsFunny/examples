/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-10 19:41 
# @File : lt_713_Subarray_Product_Less_Than_K.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	排列组合问题: 计算相乘的积小于k的组合数的总和
	双指针
	左指针的每次移动,都需要将当前的值相除直到小于目标值之后右指针继续移动
 */
//
// func numSubarrayProductLessThanK(nums []int, k int) int {
// 	if k < 0 {
// 		return 0
// 	}
//
// 	value, left, right, res := 1, 0, 0, 0
// 	for ; right < len(nums); right++ {
// 		value *= nums[right]
// 		for value >= k && left< len(nums) {
// 			value /= nums[left]
// 			left++
// 		}
// 		res+=right-left+1
// 	}
//
// 	return res
// }

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	prod := 1
	i := 0
	for j := 0; j < len(nums); j++ {
		prod *= nums[j]
		for i <= j && prod >= k {
			prod /= nums[i]
			i++
		}

		res += j-i+1
	}
	return res
}