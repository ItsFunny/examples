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
	if len(nums) == 0 {
		return nil
	}
	m := make(map[int]struct{})

}
