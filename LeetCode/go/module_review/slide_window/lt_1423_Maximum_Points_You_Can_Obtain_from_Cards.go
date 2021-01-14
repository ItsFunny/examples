/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-19 19:54 
# @File : lt_1423_Maximum_Points_You_Can_Obtain_from_Cards.go
# @Description : 
# @Attention : 
*/
package slide_window

import "math"

// 要么从左边开始拿起,要么从右边开始拿起
// 逆向思维,如果是左边或者右边k个值最大,则认为是中间连续的达到最小
func maxScore(cardPoints []int, k int) int {
	if len(cardPoints) == 0 {
		return 0
	}
	if k== len(cardPoints){
		r:=0
		for _,v:=range cardPoints{
			r+=v
		}
		return r
	}
	sum := 0
	// 左边N-k个数达到最小值才可以使得剩下的K个数为最大值
	result := math.MaxInt32
	total, have, need, left, right := 0, 0, len(cardPoints)-k, 0, 0
	for ; right < len(cardPoints); right++ {
		have++
		sum += cardPoints[right]
		for have > need {
			sum -= cardPoints[left]
			have--
			left++
		}
		if have == need {
			result = maxScoreMin(result, sum)
		}
		total += cardPoints[right]
	}
	return total - result
}

func maxScoreMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
