/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/23 9:34 下午
# @File : lt_202_判断是否是快乐数_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_isHappyStep(t *testing.T) {
	f := func(v int) int {
		ret := 0
		for v > 0 {
			ret += (v % 10) * (v % 10)
			v /= 10
		}
		return ret
	}
	fmt.Println(f(78))
}

func Test_isHappy(t *testing.T) {
	fmt.Println(isHappy(19))
}
