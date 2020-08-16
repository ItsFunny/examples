/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 13:33
# @File : select_sort.go
# @Description :
# @Attention :
*/
package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	array:=[]int{1,3,0,2,8,4,9,5}
	SelectSort(array)
	fmt.Println(1)
}
