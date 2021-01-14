/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-14 09:00
# @File : lt_81_Search_in_Rotated_Sorted_Array_II.go
# @Description :
# @Attention :
*/
package array

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	a := []int{5,1,3}
	target := 3
	b := search(a, target)
	fmt.Println(b)
}
