/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-05-04 10:38 
# @File : reflect.go
# @Description : 反射的util
*/
package utils

import (
	"fmt"
	"reflect"
	"unsafe"
)

func ConvT2InterfaceSlice(data interface{}) []interface{} {
	value, b := IsSlice(data)
	if !b {
		return nil
	}
	l := value.Len()
	res := make([]interface{}, l)
	for i := 0; i < l; i++ {
		res[i] = value.Index(i).Interface()
	}
	return res
}

// 判断是否是切片
func IsSlice(data interface{}) (reflect.Value, bool) {
	valueOf := reflect.ValueOf(data)
	if valueOf.Kind() == reflect.Slice {
		return valueOf, true
	}
	return valueOf, false
}

/*
当结构体中含有指针时，转换会导致垃圾回收的问题。
如果是 []byte 转 []T 可能会导致起始地址未对齐的问题 （[]byte 有可能从奇数位置切片）。
该转换操作可能依赖当前系统，不同类型的处理器之间有差异。
 */
 // []T 转换为[]X
func ConvT2TypeSlice(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		fmt.Sprintf("[ConvT2TypeSlice]Slice called with non-slice value of type %T", slice)
		return nil
	}
	if newSliceType.Kind() != reflect.Slice {
		fmt.Sprintf("[ConvT2TypeSlice]Slice called with non-slice type of type %T", newSliceType)
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}
