/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-19 16:52 
# @File : gc_test.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			fmt.Println(0)
		}
	}()

	wg.Wait()
	fmt.Println("end")
}
