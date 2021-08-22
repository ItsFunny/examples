/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/18 9:47 下午
# @File : lt_12_整数转罗马数字.go
# @Description :
# @Attention :
*/
package offer

// 关键: 列出所有排列组合
// 然后贪心解决即可
func intToRoman(num int) string {
	r := ""
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(ints); i++ {
		for num >=ints[i] {
			r += romans[i]
			num -= ints[i]
		}
	}

	return r
}
