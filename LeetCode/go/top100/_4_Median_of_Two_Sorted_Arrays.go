/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-06 14:15 
# @File : _4_Median_of_Two_Sorted_Arrays.go
# @Description :
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.
2个有序数组合并,返回中间值
# @Attention :
时间复杂度O(n)
则可以考虑申请额外的空间
归并排序
*/
package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n:= len(nums1), len(nums2)
	nums := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m || j < n {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				nums[k] = nums1[i]
				i++
			} else {
				nums[k] = nums2[j]
				j++
			}
		} else if i < m {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}

	return float64((nums[(m + n)>>1] + nums[((m + n - 1)>>1)])) / 2.0
}