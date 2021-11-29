/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/30 8:45 下午
# @File : lt_46_全排列.go
# @Description :
# @Attention :
*/
package offer

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	used := make([]bool, len(nums))
	permuteBack(nums, &ret, &[]int{}, 0, used)
	return ret
}
func permuteBack(nums []int, ret *[][]int, single *[]int, index int, used []bool) {
	if index == len(nums) {
		*ret = append(*ret, *single)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 开始填数
		if !used[i] {
			used[i] = true
			*single = append(*single, nums[index])
			permuteBack(nums, ret, single, index+1, used)
			used[i] = false
			v := *single
			v = v[:len(v)-1]
			single = &v
		}
	}
}
