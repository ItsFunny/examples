/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/13 9:03 上午
# @File : jz_30_连续子数组的最大和_test.go.go
# @Description :
# @Attention :
*/
package offer

import (
	"fmt"
	"testing"
)

func Test_Max(t *testing.T) {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	fmt.Println(FindGreatestSumOfSubArray(arr))
}
