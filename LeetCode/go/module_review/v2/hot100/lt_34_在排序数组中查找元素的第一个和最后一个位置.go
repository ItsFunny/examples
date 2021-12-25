/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/25 4:02 下午
# @File : lt_34_在排序数组中查找元素的第一个和最后一个位置.go
# @Description :
# @Attention :
*/
package hot100

// 关键:
// 两次查找,一次是不停的往左找,直到找到最先出现的,第二次是向右找,直到找到最后一个出现的
// 注意,我们要保存的是当前找到的值,所以会有一个额外的变量定义返回值
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	f := func(leftDirection bool) int {
		left, right := 0, len(nums)-1
		ret := -1
		for left <= right {
			mid := left + (right-left)>>1
			if target < nums[mid] {
				right = mid - 1
			} else if target > nums[mid] {
				left = mid + 1
			} else {
				// 注意,我们要保存的是当前找到的值,所以会有一个额外的变量定义返回值
				ret = mid
				// 如果是找第一个,则不停的左移动
				if leftDirection {
					right = mid - 1
				} else {
					// 说明是找最后一个,则不停的右移动
					left = mid + 1
				}
			}
		}
		return ret
	}
	ret := []int{-1, -1}
	ret[0] = f(true)
	ret[1] = f(false)
	return ret
}
