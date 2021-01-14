/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-12 08:53 
# @File : lt_78_Subsets.go
# @Description : 
# @Attention : 
*/
package array

/*
	排列组合问题
	并且不能包含重复的组合

 */
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	solve(0, nums, nil, &res)
	return res
}

func solve(index int, nums []int, cur []int, res *[][]int) {
	if index == len(nums) {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	solve(index+1, nums, cur, res)
	solve(index+1, nums, append(cur, nums[index]), res)
}