/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/3 8:47 上午
# @File : largest_rectangle_in_histogram.go
# @Description :
# @Attention :
*/
package v2

//
// 暴力法处理,
// func largestRectangleArea(heights []int) int {
// 	if len(heights) == 0 {
// 		return 0
// 	}
// 	max := 0
// 	for i := 0; i < len(heights); i++ {
// 		minHeight := heights[i]
// 		for j := i; j < len(heights); j++ {
// 			if heights[j] < minHeight {
// 				minHeight = heights[j]
// 			}
// 			tmp := minHeight * (j - i + 1)
// 			if tmp>max{
// 				max=tmp
// 			}
// 		}
// 	}
//
// 	return max
// }
//

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	max := 0
	// 关键是
	// 每当遇到下一个比当前的 下标的高度
	for i := 0; i <= len(heights); i++ {
		cur := 0
		if i != len(heights) {
			cur = heights[i]
		}
		for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
			// 开始计算之前的高度
			// 在栈中的元素,都是要比当前元素 >= 的,并且, 是必定递增的形式
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h:=heights[top]
			w := i
			if len(stack) > 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			v := h * w
			if v > max {
				max = v
			}
		}
		stack=append(stack,i)
	}

	return max
}
