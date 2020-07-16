/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-16 09:58 
# @File : _55_Jump_Game.go
# @Description :    青蛙跳格子
# @Attention : 
*/
package main

func canJump(nums []int) bool {
	prevGood := len(nums)-1
	for i := prevGood-1; i >= 0; i-- {
		if nums[i] >= prevGood-i {
			prevGood = i
		}
	}
	return prevGood == 0
}