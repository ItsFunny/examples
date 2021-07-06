/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/29 9:04 上午
# @File : jz_08_青蛙跳台阶.go
# @Description :
# @Attention :
*/
package offer

func jumpFloor(number int) int {
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	return jumpFloor(number-1) + jumpFloor(number-2)
}

func jumpFloor2(number int) int {
	if number <= 1 {
		return 1
	}
	a, b, c := 1, 1, 0
	for i := 2; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return a
}

func jumpFloor3(number int) int {
	dp := make([]int, number)
	if number==1{
		return 1
	}else if number==2{
		return 2
	}
	dp[0], dp[1] = 1, 2
	for i := 2; i < number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number-1]
}
