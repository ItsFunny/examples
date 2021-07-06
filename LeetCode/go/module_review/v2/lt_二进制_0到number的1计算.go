/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/17 9:13 上午
# @File : lt_二进制_0到number的1计算.go
# @Description :
# @Attention :
*/
package v2

func countBits(n int) []int {
	r := make([]int, 0)
	count(n,&r)
	// for i := 0; i <= n; i++ {
	// 	count(i, &r)
	// }
	return r
}
func count(n int, r *[]int) {
	if n == 0 {
		*r = append(*r, 0)
		return
	}
	count(n-1, r)
	c := 0
	for n > 0 {
		c++
		n = n & (n - 1)
	}
	*r = append(*r, c)
}
