/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/1 8:51 上午
# @File : jz_12_数的n次方.go
# @Description :
# @Attention :
*/
package offer

// 关键 2^8=2^(4*2) ,

func Power(base float64, exponent int) float64 {
	// write code here
	if exponent < 0 {
		base = 1 / base
		exponent *= -1
	}
	return pow(base, exponent)
}
func pow(value float64, count int) float64 {
	if count==0{
		return 1.0
	}
	ret:=pow(value,count>>1)
	// 判断奇偶数
	if count&1 > 0 {
		return ret*ret*value
	}
	return ret*ret
}
