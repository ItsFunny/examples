/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/12 10:15 上午
# @File : lt_16_最接近的三数之和_test.go.go
# @Description :
# @Attention :
*/
package hot100

import (
	"fmt"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	ints := make([]int, 0)
	ints = append(ints, 0, 2, 1, -3)
	fmt.Println(threeSumClosest(ints, 1))
}
