/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 15:44
# @File : quick_sort.go
# @Description :
# @Attention :
*/
package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	QuickSort(array)
	fmt.Println(array)
}
