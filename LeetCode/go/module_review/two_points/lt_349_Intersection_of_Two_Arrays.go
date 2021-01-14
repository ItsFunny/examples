/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-19 08:56 
# @File : lt_349_Intersection_of_Two_Arrays.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	计算2个数组的交集
	顺序无关,但是需要去除重复的
	除了用map还有其他更好的方法吗
	我去,只能用set来做
 */

func intersection(nums1 []int, nums2 []int) []int {

	if len(nums2) == 0 || len(nums1) == 0 {
		return nil
	}
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _, v := range nums1 {
		if _, exist := set1[v]; exist {
			continue
		}
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, exist := set2[v]; exist {
			continue
		}
		set2[v] = struct{}{}
	}
	result := make([]int, 0)
	for k, _ := range set1 {
		if _, exist := set2[k]; exist {
			result=append(result,k)
		}
	}
	return result
}
