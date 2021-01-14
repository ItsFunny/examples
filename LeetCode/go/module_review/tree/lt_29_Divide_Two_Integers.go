/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-29 13:33 
# @File : lt_29_Divide_Two_Integers.go
# @Description : 
# @Attention : 
*/
package tree

import "math"

func divide(dividend int, divisor int) int {
	reverse := false
	if divisor < 0 {
		reverse = !reverse
	}
	if dividend < 0 {
		reverse = !reverse
	}

	top, buttom := int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))

	result := 0
	last := 0
	for top >= buttom {
		last = buttom
		if top >= buttom {
			result++
			buttom <<= 1
		} else {
			top = top - last
			buttom = int(math.Abs(float64(divisor)))
		}
	}

	if reverse {
		return result * -1
	}
	return result
}
