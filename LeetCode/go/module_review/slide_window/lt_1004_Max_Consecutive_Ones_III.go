/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-16 08:59 
# @File : lt_1004_Max_Consecutive_Ones_III.go
# @Description : 
# @Attention : 
*/
package slide_window

/*
	更改后连续的最长子串
	可以认为是滑动窗口内,最多有K个0
 */

func longestOnes(A []int, K int) int {
	if len(A) == 0 {
		return 0
	}
	if K >= len(A) {
		return len(A)
	}
	max, left, right, count := 0, 0, 0, 0
	for right < len(A) {
		if A[right] == 0 {
			count++
		}
		for count > K {
			if A[left] == 0 {
				count--
			}
			left++
		}
		max = longestOnesMax(max, right-left+1)
		right++
	}

	return max
}
func longestOnesMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
