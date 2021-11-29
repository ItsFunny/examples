/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/2 8:57 下午
# @File : lt_56_合并区间.go
# @Description :
# @Attention :
*/
package offer

func merge(intervals [][]int) [][]int {
	r := make([][]int, 0)

	for i := 0; i < len(intervals); {
		j := i + 1
		max := intervals[i][1]
		for ; j < len(intervals); j++ {
			if intervals[j][0] > intervals[i][0] {
				if intervals[j][1] > max {
					max = intervals[j][1]
				}
			}
		}
		r = append(r, []int{intervals[i][0], max})
		i = i + 1
	}

	return r
}

func mergeQuickSort(intervals [][]int, left, right int) {
	if left < right {
		paration := mergeQuickSortParation(intervals, left, right)
		mergeQuickSort(intervals, left, paration)
		mergeQuickSort(intervals, paration+1, right)
	}
}
func mergeQuickSortParation(intervals [][]int, left, right int) int {
	standard := intervals[left]
	for left < right {
		for ; right > left && intervals[right][0] > standard[0]; right-- {
		}
		intervals[left] = intervals[right]
		for ; left < right && intervals[right][0] <= standard[0]; left++ {
		}
		intervals[right] = intervals[left]
	}
	intervals[left] = standard
	return left
}
