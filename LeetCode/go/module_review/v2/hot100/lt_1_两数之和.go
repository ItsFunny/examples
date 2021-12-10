/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/5 9:24 上午
# @File : lt_1_两数之和.go
# @Description :
# @Attention :
*/
package hot100

// 关键是通过一个map存储剩余的值即可,如果存在,直接返回
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, v := range nums {
		left := target - v
		if leftIndex, exist := m[left]; exist {
			return []int{index, leftIndex}
		}
		m[v] =index
	}
	return nil
}
