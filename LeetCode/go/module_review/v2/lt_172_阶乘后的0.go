/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/21 1:30 下午
# @File : lt_172_阶乘后的0.go
# @Description :
# @Attention :
*/
package v2

// 关键
// 数学规律题,背题大法好: 统计5出现的次数即可,因为为0 只有 2*5 的情况,有5必有2 ,所以统计5的次数
func trailingZeroes(n int) int {
	ret := 0
	for n > 0 {
		n /= 5
		ret += n
	}
	return ret
}
