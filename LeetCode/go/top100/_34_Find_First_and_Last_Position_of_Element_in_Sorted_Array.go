/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-24 10:14 
# @File : _34_Find_First_and_Last_Position_of_Element_in_Sorted_Array.go
# @Description :
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].
# @Attention : 
*/
package main
func searchRange(nums []int, target int) []int {
	left, right, mid := 0, len(nums)-1, 0
	for ;left<=right; {
		mid = (left+right)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if nums[left] < target {
				left ++
			}
			if nums[right] > target {
				right --
			}
			if nums[left] == nums[right] {
				return []int{left, right}
			}
		}
	}
	return []int{-1, -1}
}
