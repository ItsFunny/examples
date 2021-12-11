/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/10 8:53 下午
# @File : lt_9_判断这个数是否是回文数.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 转换为字符串,进行互相匹配即可
// func isPalindrome(x int) bool {
// 	str := strconv.Itoa(x)
// 	l := len(str)
// 	for i, j := 0, l-1; i < j; {
// 		if str[i] != str[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }

// 第二种方法,反转这个数字即可
func isPalindrome(x int) bool {
	rever := 0
	prev := x
	for x > 0 {
		rever = rever*10 + x%10
		x /= 10
	}
	return rever == prev
}
