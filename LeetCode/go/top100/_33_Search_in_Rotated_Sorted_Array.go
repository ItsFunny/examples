/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-19 14:52 
# @File : _33_Search_in_Rotated_Sorted_Array.go
# @Description :

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).
# @Attention : 
*/
package main
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	sentinel := nums[0]
	// find the rotate-pos
	// nums[rotate-pos] < sentinel && nums[rotate-pos-1] > sentinel
	l, r, rotatePos := 0, len(nums)-1, -1

	if nums[l] > nums[r] {
		for l <= r && rotatePos != (l + r) >> 1 {
			rotatePos = (l + r) >> 1
			if nums[rotatePos] >= sentinel {
				l = rotatePos+1
			} else if nums[rotatePos-1] > sentinel {
				break
			} else {
				r = rotatePos-1
			}
		}

		if target == sentinel {
			return 0
		} else if target > sentinel {
			l, r = 0, rotatePos-1
		} else {
			l, r = rotatePos, len(nums)-1
		}
	}

	targetPos := -1
	for l <= r && targetPos != (l + r) >> 1 {
		targetPos = (l + r) >> 1
		if nums[targetPos] == target {
			return targetPos
		}
		if nums[targetPos] < target {
			l = targetPos+1
		} else {
			r = targetPos-1
		}
	}

	return -1
}