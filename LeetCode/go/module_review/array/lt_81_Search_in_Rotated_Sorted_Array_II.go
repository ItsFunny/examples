/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-14 09:00 
# @File : lt_81_Search_in_Rotated_Sorted_Array_II.go
# @Description : 
# @Attention : 
*/
package array

/*
	判断值是否存在
	有序数组在某个下标进行了反转,然后判断是否存在
	有序中查找,肯定是二分最快
 */

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	halfSearch := func(nums [] int, targe int) bool {
		start, end := 0, len(nums)-1
		for start <= end {
			mid := start + (end-start)>>1
			if nums[mid] > target {
				end = mid-1
			} else if nums[mid] == target {
				return true
			} else {
				start = mid+1
			}
		}
		return false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
		// 说明在此处发生了反转,则进行二分查找
		if i-1 >= 0 && nums[i] < nums[i-1] {
			return halfSearch(nums[i:], target)
		}
	}
	return false
}
