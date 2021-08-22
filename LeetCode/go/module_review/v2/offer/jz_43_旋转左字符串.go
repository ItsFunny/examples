/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/22 8:42 上午
# @File : jz_43_旋转左字符串.go
# @Description :
# @Attention :
*/
package offer

func LeftRotateString(str string, n int) string {
	l := len(str)
	var index int
	r := make([]byte, len(str))
	for i := 0; i < l; i++ {
		index = i - n
		if index < 0 {
			index *= -1
			index = l - index
		}
		r[index] = str[i]
	}
	return string(r)
}
