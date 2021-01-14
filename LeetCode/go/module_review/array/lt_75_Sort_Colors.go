/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-09 09:11 
# @File : lt_75_Sort_Colors.go
# @Description : 
# @Attention : 
*/
package array

/*
	O(1)的复杂度排序,数字只有 0,1,2
 */
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	m := make(map[int]int)
	for _, v := range nums {
		if _, exist := m[v]; exist {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	index := 0
	for i := 0; i < 3; i++ {
		count := m[i]
		for j := 0; j < count; j++ {
			nums[index] = i
			index++
		}
	}
}
