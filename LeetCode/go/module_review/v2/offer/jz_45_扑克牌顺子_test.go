/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/23 9:48 上午
# @File : jz_45_扑克牌顺子_test.go.go
# @Description :
# @Attention :
*/
package offer

import (
	"fmt"
	"testing"
)

func TestIsContinuous(t *testing.T) {
	v := IsContinuous([]int{0, 3, 2, 6, 4})
	fmt.Println(v)
}
