/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/11 7:07 下午
# @File : lt_88_合并两个有序数组.go
# @Description :
# @Attention :
*/
package offer

// 题目关键: 数组排序递增的
// 解题关键: 双指针
func merge1(nums1 []int, m int, nums2 []int, n int) {
	for mp, np, cur := m-1, n-1, m+n-1; mp >= 0 || np >= 0; cur-- {
		var v int
		if mp == -1 {
			v = nums2[np]
			np--
		} else if np == -1 {
			v = nums1[mp]
			mp--
		} else if nums1[mp] > nums2[np] {
			v = nums1[mp]
			mp--
		} else {
			v = nums2[np]
			np--
		}
		nums1[cur] = v
	}
}
