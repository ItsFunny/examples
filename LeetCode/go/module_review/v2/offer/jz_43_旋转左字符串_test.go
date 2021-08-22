/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/22 8:47 上午
# @File : jz_43_旋转左字符串_test.go.go
# @Description :
# @Attention :
*/
package offer

import (
	"fmt"
	"testing"
)

func TestLeftRotateString(t *testing.T) {
	v := LeftRotateString("abcXYZdef", 3)
	fmt.Println(v)
}
