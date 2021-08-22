/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/15 4:25 下午
# @File : lt_4_有序数组的中位数.go
# @Description :
# @Attention :
*/
package offer

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	left, right := 0, 0
	start1, start2 := 0, 0
	for i := 0; i <= (l1+l2)>>1; i++ {
		left = right
		if start1 < l1 && (nums1[start1] < nums2[start2] || start2 >= l2) {
			right = nums1[start1]
			start1++
		} else {
			right = nums2[start2]
			start2++
		}
	}
	if (l1+l2)&1 == 0 {
		return float64(left+right) / 2.0
	}
	return float64(right)
}
