/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-12 10:08 
# @File : list.go
# @Description : 
# @Attention : 
*/
package array_list

type List interface {
	Add(data interface{})
	RemoveByIndex(index int) interface{}
	Show() func() (interface{}, bool)
	Size()int
}
