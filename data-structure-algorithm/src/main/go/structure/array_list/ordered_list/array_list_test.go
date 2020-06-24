/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-10 12:50 
# @File : array_list_test.go
# @Description : 
# @Attention : 
*/
package ordered_list

import (
	"fmt"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	list := NewArrayList()
	list.Add(1)
	list.Add(2)

	iterator := list.Iterator()
	for {
		data, b := iterator()
		if !b {
			break
		}
		fmt.Println(data)
	}

	list.RemoveByIndex(0)
	iterator = list.Iterator()
	for {
		data, b := iterator()
		if !b {
			break
		}
		fmt.Println(data)
	}

}
