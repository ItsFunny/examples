/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/10 6:36 下午
# @File : lt_8_字符串转换整数atoi_test.go.go
# @Description :
# @Attention :
*/
package hot100

import (
	"fmt"
	"testing"
	"time"
)

func Test_myAtoi(t *testing.T) {
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(1<<31 - 1)
}

type WatchMessage interface {
	GetKey() []byte
	GetValue() string
	GetType() uint32
}

func TestA(t *testing.T) {
	ret := []WatchMessage{}
	ret = append(ret, nil)
	for _, v := range ret {
		fmt.Println(v.GetValue(), v.GetKey(), v.GetType())
	}
}

func TestRouting(t *testing.T) {
	c := make(chan error, 1)
	go func() {
		for {
			select {
			case v, ok := <-c:
				fmt.Println(v, ok)
				if !ok {
					fmt.Println("done")
					return
				}
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)
		c <- nil
		time.Sleep(time.Second * 3)
		close(c)
	}()
	select {}
}

func TestPrint(t *testing.T){
	//9223372036854775708
	//1000000000000000
	fmt.Println(1<<63-100)
}