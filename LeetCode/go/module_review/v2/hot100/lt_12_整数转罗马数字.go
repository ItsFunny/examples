/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/11 8:56 上午
# @File : lt_12_整数转罗马数字.go
# @Description :
# @Attention :
*/
package hot100

// func intToRoman(num int) string {
// 	r := ""
// 	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
// 	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
//
// 	for i := 0; i < len(ints); i++ {
// 		for num >= ints[i] {
// 			r += romans[i]
// 			num -= ints[i]
// 		}
// 	}
//
// 	return r
// }

// 关键: 先列举出所有 数字对应的罗马数字,然后遍历 ,如1000 , num有几个1000,就可以加几个1000对应的罗马数字即可
func intToRoman(num int) string {
	ret := ""
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, k := range nums {
		if num == 0 {
			break
		}
		v := m[k]
		if num/k > 0 {
			count := num / k
			for i := 0; i < count; i++ {
				ret += v
			}
			num %= k
		}
	}

	return ret
}
