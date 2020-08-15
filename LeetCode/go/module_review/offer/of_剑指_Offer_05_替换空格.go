/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-14 10:39 
# @File : of_剑指_Offer_05_替换空格.go
# @Description : 
# @Attention : 
*/
package offer

func replaceSpace(s string) string {
	bytes := make([]byte, 0)
	for i, j := 0, 0; i < len(s); {
		if s[i] == ' ' {
			bytes = append(bytes, '%', '2', '0')
			// bytes[j] = '%'
			// bytes[j+1] = '2'
			// bytes[j+2] = '0'
			j += 3
		} else {
			bytes = append(bytes, s[i])
			j++
		}
		i++
	}
	return string(bytes)
}
