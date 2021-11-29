/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/9 8:44 上午
# @File : lt_78_子集.go
# @Description :
# @Attention :
*/
package offer

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	subsetsBacktrack(nums, 0, &ret)
	return ret
}

var current []int

func subsetsBacktrack(nums []int, index int, ret *[][]int) {
	if index == len(nums) {
		*ret = append(*ret, append([]int(nil), current...))
		return
	}
	current = append(current, nums[index])
	// for i := index; i < len(nums); i++ {
	subsetsBacktrack(nums, index+1, ret)
	current = current[:len(current)-1]
	subsetsBacktrack(nums, index+1, ret)
	// }
}
