/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-08 10:21 
# @File : array_lisr.go
# @Description :   顺序表,本质为数组
# @Attention : 
*/
package ordered_list

import (
	"errors"
	"math"
)

type arrayList struct {
	data []interface{}
	// 长度
	size int
	//
	modCount int
}

func NewArrayList() *arrayList {
	list := new(arrayList)
	return list
}

func (this *arrayList) Add(data interface{}) {
	this.ensureCapacityInternal(this.calcCapacity(this.size + 1))
	this.data[this.size] = data
	this.size++
}

func (this *arrayList) ensureCapacityInternal(minCap int) {
	this.modCount++

	if minCap-len(this.data) > 0 {
		// 扩容
		this.grow(minCap)
	}
}

func (this *arrayList) calcCapacity(minCap int) int {
	if this.data == nil {
		return int(math.Max(10, float64(minCap)))
	}
	return minCap
}

func (this *arrayList) grow(cap int) {
	oldCap := len(this.data)
	newCap := oldCap + (oldCap << 1)
	if newCap-cap < 0 {
		newCap = cap
	}
	// 切片复制
	newData := make([]interface{}, newCap)
	copy(this.data, newData[:len(this.data)])
	this.data = newData
	// chou ni die ne
}

// 通过元素下标删除
func (this *arrayList) RemoveByIndex(index int) error{
	if e:=this.rangeCheck(index);nil!=e {
		return e
	}
	// copy()
	this.data=append(this.data[:index],this.data[index+1:])

	return nil
}
func (this *arrayList) rangeCheck(index int) error {
	if index < 0 || index > this.size {
		return errors.New("数组下标越界错误")
	}
	return nil
}

// func (this *arrayList) Iterator() (int, bool) {
// 	index := 0
// 	return func() (value int, exist bool) {
// 		return value, exist
// 	}
// }

func (this *arrayList) Iterator() func() (interface{}, bool) {
	index := 0
	return func() (val interface{}, ok bool) {
		if index >= len(this.data) {
			return
		}
		val, ok = this.data[index], true
		index++
		return
	}
}
