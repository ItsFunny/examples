/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-13 17:16 
# @File : slice.go
# @Description : 切片util
*/
package utils

func ClearPointerSlice(slices ...*[]interface{}) {
	for _, s := range slices {
		*s = (*s)[0:0]
	}
}