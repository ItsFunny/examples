/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-16 14:07
# @File : shell_sort.go
# @Description :
# @Attention :
*/
package sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	ShellSort(array)
	fmt.Println(array)
}
