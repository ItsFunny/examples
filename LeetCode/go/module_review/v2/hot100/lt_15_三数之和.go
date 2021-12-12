/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/12 9:17 上午
# @File : lt_15_三数之和.go
# @Description :
# @Attention :
*/
package hot100

import "sort"

// 判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	//  第一步: 先排序
	sort.Ints(nums)
	// 第二步: 选择基准值,然后进行for循环遍历
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		// 第三步: a去除重复的解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 第三步: 首尾双指针
		for j, k := i+1, len(nums)-1; j < k; {
			b, c := nums[j], nums[k]
			if a+b+c < 0 {
				j++
			} else if a+b+c > 0 {
				k--
			} else {
				ret = append(ret, []int{a, b, c})
				// 第四步: 去除b,c重复解,就是将 等于 b,c的值全都过滤掉
				j++
				k--
				for ; j < k && b == nums[j]; j++ {}
				for ; k > j && c == nums[k]; k-- {}
			}
		}
	}
	return ret
}
