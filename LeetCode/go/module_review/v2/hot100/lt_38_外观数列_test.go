/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/28 9:13 上午
# @File : lt_38_外观数列_test.go.go
# @Description :
# @Attention :
*/
package hot100

import (
	"fmt"
	"testing"
	"time"
)

func Test_countAndSay(t *testing.T) {
	fmt.Println(countAndSay(4))
}

func Test_Sleep(t *testing.T) {
	ars := []int{1, 2, 3}
	for _, v := range ars {
		go func() {
			fmt.Println(v)
		}()
		time.Sleep(time.Second)
	}
}

func TestClosure(t *testing.T) {
	addrList := []string{
		"上地",
		"马连洼",
		"五道口",
		"西二旗",
	}
	for _, addr := range addrList {
		//v := addr
		test := func(addrPassedByValue string) {
			testFunc1(addrPassedByValue)
			testFunc2(addr)
		}

		go runTest(test, addr)
		time.Sleep(time.Millisecond)
	}

	time.Sleep(5 * time.Second)
}

func runTest(f func(string), arg string) {
	f(arg)
}

func testFunc1(addrPassedByValue string) {
	fmt.Printf("testFunc1 addrPassedByValue: %s\n", addrPassedByValue)
}
func testFunc2(addr string) {
	fmt.Printf("testFunc2 addr: %s\n", addr)
}

func TestEcho(t *testing.T) {
	f := func() {
		fmt.Println(1)
		defer func() { fmt.Println(2) }()
		defer func() { fmt.Println(3) }()
	}
	f()
}

type A struct {
	name string
}

func TestEEE(t *testing.T) {
	a := &A{name: "charlie"}
	f := func(v **A) {
		b:=&A{name: "joker"}
		*v=b
	}
	f(&a)
	fmt.Println(a.name)
}
