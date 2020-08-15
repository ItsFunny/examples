/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-15 12:58 
# @File : lt_153_Find_Minimum_in_Rotated_Sorted_Array.go
# @Description :
假设按照升序排序的数组在预先未知的某个点上进行了旋转
( 例如，数组  [0,1,2,4,5,6,7] 可能变为  [4,5,6,7,0,1,2] )。 请找出其中最小的元素。
# @Attention :
	因为发生了反转,所以肯定不是升序的,意味着当左边的>右边的就不是正常的
*/
package half

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)>>1
		//
		if nums[mid] > nums[end] {
			start=mid
		} else {
			end = mid
		}
	}
	if nums[start]<nums[end] {
		return nums[start]
	}
	return nums[end]
}
