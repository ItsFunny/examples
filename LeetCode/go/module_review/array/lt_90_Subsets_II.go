/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-17 14:44 
# @File : lt_90_Subsets_II.go
# @Description : 
# @Attention : 
*/
package array

import "sort"

/*
	求元素的子集
 */

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)

	result := make([][]int, 0)
	subsetsWithDupBackTrack(nums, &result, []int{})
	return result
}

func subsetsWithDupBackTrack(nums []int, result *[][]int, temp []int) {
	*result = append(*result, append([]int{}, temp...))

	for i:=0;i< len(nums);i++{
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		subsetsWithDupBackTrack(nums[i+1:],result,append(temp,nums[i]))
	}

}
