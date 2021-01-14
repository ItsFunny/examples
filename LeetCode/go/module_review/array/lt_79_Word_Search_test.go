/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-13 08:43
# @File : lt_79_Word_Search.go
# @Description :
# @Attention :
*/
package array

import (
	"fmt"
	"testing"
)

func Test_exist(t *testing.T) {
	// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	str := "ABCCED"
	b := exist(board, str)
	fmt.Println(b)
}
