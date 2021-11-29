/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/23 9:15 下午
# @File : lt_202_判断是否是快乐数.go
# @Description :
# @Attention :
*/
package v2

// 关键: 判断算出来的数,是否在map中存在,存在则代表着有环,必然无法结果为1
func isHappy(n int) bool {
	m := make(map[int]struct{})
	for n != 1 {
		n = isHappyStep(n)
		if _, exist := m[n]; exist {
			return false
		}
		m[n] = struct{}{}
	}
	return n == 1
}
func isHappyStep(v int) int {
	ret := 0
	for v > 0 {
		ret += (v % 10) * (v % 10)
		v /= 10
	}
	return ret
}
