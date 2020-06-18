/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 09:55 
# @File : linked_list_test.go
# @Description : 
# @Attention : 
*/
package linked_list

import "testing"

func TestNewLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Add(1)
	linkedList.Show()
	linkedList.Add(2)
	linkedList.Show()
}

func TestLinkedList_RemoveByIndex(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Add(1)
	linkedList.Show()
	linkedList.Add(2)
	linkedList.Show()

	linkedList.RemoveByIndex(0)
	linkedList.Show()

}
