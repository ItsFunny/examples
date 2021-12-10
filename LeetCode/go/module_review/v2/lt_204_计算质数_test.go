/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/25 8:46 下午
# @File : lt_204_计算质数_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_countPrimes(t *testing.T) {
	fmt.Println(countPrimes(10))
}

func TestPrint(t *testing.T) {
	for i := 1; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
		// time.Sleep(time.Second)
	}
	fmt.Println("end")
}

func TestString(t *testing.T) {
	str := "asddd"
	bs := []byte(str)
	bs[1] = 'd'
	fmt.Println(str)
}

type B struct {
	name string
}
type A struct {
	b *B
}

func TestMap(t *testing.T) {
	m := make(map[int]A)
	m[1] = A{b: &B{name: "joker"}}
	a := m[1]
	a.b.name = "charlie"
	fmt.Println(m[1].b.name)
}
