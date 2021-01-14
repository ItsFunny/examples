/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-16 09:45
# @File : lt_1052_Grumpy_Bookstore_Owner.go
# @Description :
# @Attention :
*/
package slide_window

import (
	"fmt"
	"testing"
)

func Test_maxSatisfied(t *testing.T) {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	   grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	X := 3
	satisfied := maxSatisfied(customers, grumpy, X)
	fmt.Println(satisfied)
}
