/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-16 08:59
# @File : lt_1004_Max_Consecutive_Ones_III.go
# @Description :
# @Attention :
*/
package slide_window

import (
	"fmt"
	"testing"
)

func Test_longestOnes(t *testing.T) {
	A := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K := 2
	ones := longestOnes(A, K)
	fmt.Println(ones)

}
