/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 15:56
# @File : heap_sort.go
# @Description :
# @Attention :
*/
package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	HeapSort(array)
	fmt.Println(array)
}
