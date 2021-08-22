/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/29 8:46 上午
# @File : lc_1_两数和.go
# @Description :
# @Attention :
*/
package offer

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, v := range nums {
		vv, exist := m[target-v]
		if exist {
			return []int{vv, index}
		}
		m[v] = index
	}
	return nil
}
