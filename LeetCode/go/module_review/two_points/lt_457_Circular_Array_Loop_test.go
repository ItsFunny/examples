/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-09 09:50
# @File : lt_457_Circular_Array_Loop.go
# @Description :
# @Attention :
*/
package two_points

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	nums:=[]int{-1,2,1,2}
	circularArrayLoop(nums)
}

func Test_setZero(t *testing.T) {
	type args struct {
		nums []int
		i    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZero(tt.args.nums, tt.args.i)
		})
	}
}
