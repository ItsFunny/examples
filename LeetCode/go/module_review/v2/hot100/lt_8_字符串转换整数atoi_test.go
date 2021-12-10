/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/10 6:36 下午
# @File : lt_8_字符串转换整数atoi_test.go.go
# @Description :
# @Attention :
*/
package hot100

import (
	"fmt"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(1<<31 - 1)
}
