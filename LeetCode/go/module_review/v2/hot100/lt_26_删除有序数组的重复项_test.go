/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/17 9:15 上午
# @File : lt_26_删除有序数组的重复项_test.go.go
# @Description :
# @Attention :
*/
package hot100

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 2, 3}))
}
func TestCloseCh(t *testing.T){
	ch:=make(chan struct{},1)
	v,ok:=<-ch
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}