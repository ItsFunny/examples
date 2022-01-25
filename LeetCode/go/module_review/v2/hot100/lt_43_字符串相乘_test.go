/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/12 9:10 上午
# @File : lt_43_字符串相乘_test.go.go
# @Description :
# @Attention :
*/
package hot100

import (
	"crypto/md5"
	"fmt"
	"math"
	"testing"
	"time"
)

func Test_asd(t *testing.T) {
	h1 := make([]int, 0)
	for i := 0; i < 100000; i++ {
		go func() {
			for i := 0; i <  math.MaxInt64; i++ {
				h1 = append(h1, i)
			}
		}()
		go func() {
			for i := 0; i <  math.MaxInt64; i++ {
				h1 = append(h1, i)
			}
		}()
	}

	time.Sleep(time.Second * 100)
}

func Test_Arrs(t *testing.T) {
	data := []byte("asd")
	ret := md5.Sum(data)
	m := make(map[string]struct{})
	r1 := fmt.Sprintf("%v-%d", ret, 1)
	r2 := fmt.Sprintf("%v-%d", ret, 1)
	_, exist := m[r1]
	fmt.Println(exist)
	if !exist {
		m[r1] = struct{}{}
	}
	_, exist = m[r2]
	fmt.Println(exist)
	if !exist {
		m[r2] = struct{}{}
	}
}

// go:linkname memmove runtime.memmove

//func TestMemmove(t *testing.T)  {
//	bytes := []byte("test")
//	bytes1 := []byte("test")
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println(r)
//		}
//	}()
//	wg := sync.WaitGroup{}
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func() {
//			memmove(unsafe.Pointer(&bytes), unsafe.Pointer(&bytes1), 1024000000)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//}

func Test_multiply(t *testing.T) {
	fmt.Println(multiply("9","9"))
}