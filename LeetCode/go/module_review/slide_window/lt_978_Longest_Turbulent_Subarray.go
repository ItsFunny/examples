/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-13 09:32 
# @File : lt_978_Longest_Turbulent_Subarray.go
# @Description : 
# @Attention : 
*/
package slide_window

// func maxTurbulenceSize(A []int) int {
// 	if len(A) == 0 {
// 		return 0
// 	}
// 	after := make([]int, 0)
// 	for i, j := 0, 1; j < len(A); j++ {
// 		if A[i] > A[j] {
// 			after = append(after, 1)
// 		} else if A[i] == A[j] {
// 			j++
// 			after = append(after, 0)
// 		} else {
// 			after = append(after, -1)
// 		}
// 		i++
// 	}
// 	after = append(after, 1)
// 	return maxTurbulenceSize2(after)
// }
//
// // 最长交替子序列
// func maxTurbulenceSize2(nums []int) int {
// 	max := 0
// 	left, right := 0, 1
// 	if len(nums) == 1 {
// 		return 1
// 	}
// 	for ; right < len(nums); right++ {
// 		// 判断是否前后交替: 说明不是前后交替,则需要更新最大值
// 		if c := nums[right-1] * nums[right]; c != -1 {
// 			if c != 0 {
// 				max = maxTurbulenceSize2Max(max, right-left+1)
// 			}
// 			left = right
// 		}
// 	}
// 	return max
// }
//
// func maxTurbulenceSize2Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func maxTurbulenceSize(A []int) int {
	if len(A) == 0 {
		return 0
	}
	max := 1
	left, right := 0, 1
	for ; right < len(A); right++ {
		compare := maxTurbulenceSizeCompare(A[right-1], A[right])
		// 说明不是交替了,而是两个相同符号
		if right == len(A)-1 || compare*maxTurbulenceSizeCompare(A[right], A[right+1]) != -1 {
			if compare != 0 {
				// 需要计算最大值
				max = maxTurbulenceSizeMax(max, right-left+1)
			}
			// 说明的是上一个元素与这个元素相同,依然左指针需要进行移动
			left = right
		}
	}

	return max
}

func maxTurbulenceSizeCompare(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func maxTurbulenceSizeMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
