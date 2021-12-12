/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/12 10:06 上午
# @File : lt_16_最接近的三数之和.go
# @Description :
# @Attention :
*/
package hot100

import (
	"math"
	"sort"
)

// 关键: 排序+双指针
func threeSumClosest(nums []int, target int) int {
	ret := math.MaxInt32
	// 第一步: 排序
	sort.Ints(nums)
	var a, b, c int
	update := func(cur int) {
		r1 := cur - target
		r2 := ret - target
		if r1 < 0 {
			r1 *= -1
		}
		if r2 < 0 {
			r2 *= -1
		}
		if r1 < r2 {
			ret = cur
		}
	}
	for i := 0; i < len(nums); i++ {
		// 第二步: 以a 为基准值
		a = nums[i]
		// 第三步: 去除a的重复解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			// 第四步: 定b,c的基准
			b, c = nums[j], nums[k]
			sum := a + b + c
			if sum == target {
				return target
				// 因为数组是已经排序过了的,所以可以控制移动方向
			} else if sum > target {
				// 如果比target大,左移,并且去除 c 重复解
				k--
				for ; k > j && c == nums[k]; k-- {
				}
			} else {
				// 如果比target小,右移,去除b重复解
				j++
				for ; j < k && nums[j] == b; j++ {
				}
			}
			// 第五步: 更新返回值
			update(sum)
		}
	}
	return ret
}
