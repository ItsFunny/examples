/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/1 9:19 上午
# @File : lt_242_有效的字母异位词.go
# @Description :
# @Attention :
*/
package v2

// 关键: 出现的次数都要相等,代表着 排序之后,结果要一致
// 或者是通过hash表的方式,因为字母只有26个
func isAnagram(s string, t string) bool {
	var c1 [26]byte
	var c2 [26]byte
	for _, v := range s {
		c1[v-'a']++
	}
	for _, v := range t {
		c2[v-'a']++
	}
	return c1 == c2
}
