/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-11 10:53 
# @File : _11_Container_With_Most_Water.go
# @Description :   最多装水量
# @Attention : 
*/
package main

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		res = max(res, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}