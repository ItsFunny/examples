/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-02-01 17:35 
# @File : slice.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"errors"
	"reflect"
)

// 删除切片
func DeleteSlice(slice interface{}, index int) (interface{}, error) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	if slice == nil || length == 0 || (length-1) < index {
		return nil, errors.New("error")
	}
	if length-1 == index {
		return sliceValue.Slice(0, index).Interface(), nil
	} else if (length - 1) >= index {
		return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface(), nil
	}
	return nil, errors.New("error")
}
