/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-23 09:23
# @File : lt_57_Insert_Interval.go
# @Description :
# @Attention :
*/
package array

import (
	"testing"
)

func Test_insert(t *testing.T) {
	a:=make([][]int,0)
	// {
	// 	[]int{1,3},
	// 		[]int{6,9},
	// }
	a=append(a,[]int{1,3})
	a=append(a,[]int{6,9})
	insert(a,[]int{2,5})
}

func Test_insertQs(t *testing.T) {
	type args struct {
		ints  [][]int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertQs(tt.args.ints, tt.args.left, tt.args.right)
		})
	}
}

func Test_insertParation(t *testing.T) {
	type args struct {
		ints  [][]int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertParation(tt.args.ints, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("insertParation() = %v, want %v", got, tt.want)
			}
		})
	}
}
