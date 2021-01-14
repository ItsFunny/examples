/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-19 09:09 
# @File : lt_350_Intersection_of_Two_Arrays_II.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	2个数组的交集,但是不去重,顺序无关
 */
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) == 0 || len(nums1) == 0 {
		return nil
	}
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v := range nums1 {
		if _, exist := m1[v]; exist {
			m1[v]++
		} else {
			m1[v] = 1
		}
	}
	for _, v := range nums2 {
		if _, exist := m2[v]; exist {
			m2[v]++
		} else {
			m2[v] = 1
		}
	}

	result := make([]int, 0)
	c := 0
	for k, v1 := range m1 {
		if v2, exist := m2[k]; exist {
			if v1 < v2 {
				c = v1
			} else {
				c = v2
			}
			for i := 0; i < c; i++ {
				result = append(result, k)
			}
		}
	}
	return result
}
