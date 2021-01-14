/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-17 09:32
# @File : lt_76_Minimum_Window_Substring.go
# @Description :
# @Attention :
*/
package slide_window

import (
	"fmt"
	"testing"
)

func Test_minWindow(t *testing.T) {
	window := minWindow2("ADOBECODEBANC", "ABC")
	fmt.Println(window)
}
