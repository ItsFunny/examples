/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-14 08:26
# @File : lt_80_Remove_Duplicates_from_Sorted_Array_II.go
# @Description :
# @Attention :
*/
package array

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	a := []int{1, 1, 1, 2, 2, 3}
	count := removeDuplicates(a)
	fmt.Println(count)
}
