/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-19 09:16
# @File : lt_350_Intersection_of_Two_Arrays_II_test.go.go
# @Description :
# @Attention :
*/
package two_points

import (
	"fmt"
	"testing"
)


func Test_intersect(t *testing.T) {
	n1 := []int{4, 9, 5}
	n2 := []int{9, 4, 9, 8, 4}
	ints := intersect(n1, n2)
	fmt.Println(ints)
}
