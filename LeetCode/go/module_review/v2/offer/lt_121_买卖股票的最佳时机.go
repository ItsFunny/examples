/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/27 9:03 上午
# @File : lt_121_买卖股票的最佳时机.go
# @Description :
# @Attention :
*/
package offer

import "math"

// 一次遍历
// 找到价格最小值,然后每次都计算利润
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := math.MaxInt32
	ret := 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v-minPrice > ret {
			ret = v - minPrice
		}
	}

	return ret
}
