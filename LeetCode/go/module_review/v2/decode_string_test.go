/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/1 8:54 上午
# @File : decode_string_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_decodeString(t *testing.T) {
	fmt.Println(decodeString("3[a]2[bc]"))
}
