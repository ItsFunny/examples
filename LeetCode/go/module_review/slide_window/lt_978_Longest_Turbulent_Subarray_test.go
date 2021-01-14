/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-13 09:32
# @File : lt_978_Longest_Turbulent_Subarray.go
# @Description :
# @Attention :
*/
package slide_window

import (
	"fmt"
	"testing"
)

func Test_maxTurbulenceSize(t *testing.T) {
	// A := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	A := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	size := maxTurbulenceSize(A)
	fmt.Println(size)
}
