/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-23 09:23 
# @File : lt_57_Insert_Interval.go
# @Description : 
# @Attention : 
*/
package array

/*
	将这个新的interval 插入,然后返回的是一个merge 过的区间
	核心:
	先插入,然后排序,再然后合并区间
 */

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval)==0{
		return intervals
	}
	if len(intervals)   ==0 || len(intervals[0])==0{
		return [][]int{newInterval}
	}
	v := make([][]int, 3)
	for i := 0; i < len(intervals); i++ {
		v[i] = intervals[i]
	}
	v[len(v)-1] = newInterval
	intervals = v

	insertQs(intervals, 0, len(intervals)-1)

	// 合并
	result := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		node := intervals[i]
		if len(result) == 0 {
			result = append(result, node)
			continue
		}
		last := result[len(result)-1]

		if node[0] <= last[len(last)-1] {
			if node[0] < last[0] {
				last[0] = node[0]
			}
			if node[len(node)-1] > last[len(last)-1] {
				last[len(last)-1] = node[len(node)-1]
			}
		} else {
			result = append(result, node)
		}
	}

	return result
}
func insertQs(ints [][]int, left, right int) {
	if left < right {
		mid := insertParation(ints, left, right)
		insertQs(ints, left, mid)
		insertQs(ints, mid+1, right)
	}
}
func insertParation(ints [][]int, left, right int) int {
	standard := ints[left]
	for left < right {
		for right > left && ints[right][0] >= standard[0] {
			right--
		}
		ints[left] = ints[right]

		for left < right && ints[left][0] <= standard[0] {
			left++
		}
		ints[right] = ints[left]
	}
	ints[left] = standard
	return left
}
