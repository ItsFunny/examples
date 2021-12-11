/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/11 3:14 下午
# @File : lt_13_罗马数字转整数.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 如果说通过当前值获取得到罗马之后,小的在前面,则 返回值需要减去,否则的话,直接累加
func romanToInt(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		v := getInt(s[i])
		if i+1 < len(s) && v < getInt(s[i+1]) {
			// 如果说小的在前面,如 IV ,I < V  ,CD C<D
			// 则此时是要减去的
			ret -= v
		} else {
			ret += v
		}
	}

	return ret
}

func getInt(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
