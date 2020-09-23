/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-21 19:59 
# @File : lt_56_Merge_Intervals.go
# @Description : 
# @Attention : 
*/
package array

/*
	合并区间:
	输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
	输出: [[1,6],[8,10],[15,18]]
	解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
	核心是排序,然后比较: 如果 某一块的左边在之前的那块内,代表可以合并(既 nodeLeft< 最后一个节点的右边界)
 */

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	if len(intervals[0]) == 0 {
		return nil
	}
	qs(intervals, 0, len(intervals)-1)
	result := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		// 判断左节点
		if len(result) == 0 {
			result = append(result, intervals[i])
			continue
		}
		// 左孩子小于最后一个右孩子的最大值,说明有重复,则直接归并掉
		last := result[len(result)-1]
		node := intervals[i]
		if node[0] <=last[len(last)-1] {
			// 直接合道最后一个元素
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

// 排序

// 快排
func qs(intetvals [][]int, left, right int) {
	if left < right {
		mid := qsparation(intetvals, left, right)
		qs(intetvals, left, mid)
		qs(intetvals, mid+1, right)
	}
}

func qsparation(ints [][]int, left int, right int) int {
	standard := ints[left]
	for left < right {
		for right > left && ints[right][0] >= standard[0] {
			right--
		}
		ints[left] = ints[right]
		for left < right && ints[left][0] <=standard[0] {
			left++
		}
		ints[right] = ints[left]
	}
	ints[left] = standard

	return left
}
