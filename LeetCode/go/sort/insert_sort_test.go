/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 13:40
# @File : insert_sort.go
# @Description : 插入排序
# @Attention :
*/
package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	InsertSort(array)
	fmt.Println(array)
}
