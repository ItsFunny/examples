/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/18 9:19 上午
# @File : sort_test.go
# @Description :
# @Attention :
*/
package sort

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	arr         = []int{0, 3, 2, 1, 9, 8}
	exceptedRet = []int{0, 1, 2, 3, 8, 9}
)

func TestBuble(t *testing.T) {
	ret := bubbleSort(arr)
	fmt.Println(ret)
}
func TestSelect(t *testing.T) {
	ret := SelectionSort(arr)
	fmt.Println(ret)
	require.Equal(t, ret, exceptedRet)
}

func TestInsert(t *testing.T) {
	ret := InsertionSort(arr)
	fmt.Println(arr)
	require.Equal(t, ret, exceptedRet)
}

func TestMerge(t *testing.T) {
	ret := mergeSort(arr)
	fmt.Println(ret)
	require.Equal(t, ret, exceptedRet)
}
func TestQSort(t *testing.T) {
	ret := quickSort([]int{1,2,3,1})
	fmt.Println(ret)
	require.Equal(t, ret, exceptedRet)
}
func TestShell(t *testing.T) {
	ret := shellSort(arr)
	fmt.Println(ret)
	require.Equal(t, ret, exceptedRet)
}
