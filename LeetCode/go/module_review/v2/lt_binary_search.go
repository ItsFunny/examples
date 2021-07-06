/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/21 9:05 上午
# @File : lt_binary_search.go
# @Description :
# @Attention :
*/
package v2

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func search2(nums []int,target int)int{
	left:=0
	right:= len(nums)-1
	for left+1<right{
		mid:=left+(right-left)>>1
		if nums[mid]==target{
			right=mid
		}else if nums[mid]<target{
			left=mid
		}else{
			right=mid
		}
	}
	if nums[right]==target{
		return right
	}
	if nums[left]==target{
		return left
	}
	return -1
}
