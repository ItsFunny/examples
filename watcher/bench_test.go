/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/22 1:01 下午
# @File : bench.go
# @Description :
# @Attention :
*/
package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

var funcs = []struct {
	name string
	f    func(...<-chan int) <-chan int
}{
	{"goroutines", goroutine},
	{"goroutineMerge", mergeN},
	{"reflection", mergeReflect},
	{"selectn", selectn},
}

func goroutine(chans ...<-chan int) <-chan int {
	r := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	go func() {
		for i := 0; i < len(chans); i++ {
			go func(index int) {
				defer wg.Done()
				for v := range chans[index] {
					r <- v
				}
			}(i)
		}
		wg.Wait()
		close(r)
	}()
	return r
}

func selectn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		i := 0
		l := len(chs)
		max := 32
		wg := sync.WaitGroup{}
		for i < len(chs) {
			l = len(chs) - i
			switch {
			case l > 32 && max >= 32:
				wg.Add(1)
				go select32(chs[i:i+32], out, &wg)
				i += 32
			case l > 15 && max >= 16:
				wg.Add(1)
				go select16(chs[i:i+16], out, &wg)
				i += 16
			case l > 7 && max >= 8:
				wg.Add(1)
				go select8(chs[i:i+8], out, &wg)
				i += 8
			case l > 3 && max >= 4:
				wg.Add(1)
				go select4(chs[i:i+4], out, &wg)
				i += 4
			case l > 1 && max >= 2:
				wg.Add(1)
				go select2(chs[i:i+2], out, &wg)
				i += 2
			case l > 0:
				wg.Add(1)
				go select1([]<-chan int{chs[i]}, out, &wg)
				i += 1
			}
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func select32(cs []<-chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var v int
	var ok bool
	done := 0
	for {
		select {
		case v, ok = <-cs[0]:
			if !ok {
				done++
				cs[0] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[1]:
			if !ok {
				done++
				cs[1] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[2]:
			if !ok {
				done++
				cs[2] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[3]:
			if !ok {
				done++
				cs[3] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[4]:
			if !ok {
				done++
				cs[4] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[5]:
			if !ok {
				done++
				cs[5] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[6]:
			if !ok {
				done++
				cs[6] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[7]:
			if !ok {
				done++
				cs[7] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[8]:
			if !ok {
				done++
				cs[8] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[9]:
			if !ok {
				done++
				cs[9] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[10]:
			if !ok {
				done++
				cs[10] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[11]:
			if !ok {
				done++
				cs[11] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[12]:
			if !ok {
				done++
				cs[12] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[13]:
			if !ok {
				done++
				cs[13] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[14]:
			if !ok {
				done++
				cs[14] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[15]:
			if !ok {
				done++
				cs[15] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[16]:
			if !ok {
				done++
				cs[16] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[17]:
			if !ok {
				done++
				cs[17] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[18]:
			if !ok {
				done++
				cs[18] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[19]:
			if !ok {
				done++
				cs[19] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[20]:
			if !ok {
				done++
				cs[20] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[21]:
			if !ok {
				done++
				cs[21] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[22]:
			if !ok {
				done++
				cs[22] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[23]:
			if !ok {
				done++
				cs[23] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[24]:
			if !ok {
				done++
				cs[24] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[25]:
			if !ok {
				done++
				cs[25] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[26]:
			if !ok {
				done++
				cs[26] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[27]:
			if !ok {
				done++
				cs[27] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[28]:
			if !ok {
				done++
				cs[28] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[29]:
			if !ok {
				done++
				cs[29] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[30]:
			if !ok {
				done++
				cs[30] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[31]:
			if !ok {
				done++
				cs[31] = nil
				if done == 32 {
					return
				}
				continue
			}
			out <- v
		}
	}

}
func select16(cs []<-chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var v int
	var ok bool
	done := 0
	for {
		select {
		case v, ok = <-cs[0]:
			if !ok {
				done++
				cs[0] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[1]:
			if !ok {
				done++
				cs[1] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[2]:
			if !ok {
				done++
				cs[2] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[3]:
			if !ok {
				done++
				cs[3] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[4]:
			if !ok {
				done++
				cs[4] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[5]:
			if !ok {
				done++
				cs[5] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[6]:
			if !ok {
				done++
				cs[6] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[7]:
			if !ok {
				done++
				cs[7] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[8]:
			if !ok {
				done++
				cs[8] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[9]:
			if !ok {
				done++
				cs[9] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[10]:
			if !ok {
				done++
				cs[10] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[11]:
			if !ok {
				done++
				cs[11] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[12]:
			if !ok {
				done++
				cs[12] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[13]:
			if !ok {
				done++
				cs[13] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[14]:
			if !ok {
				done++
				cs[14] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[15]:
			if !ok {
				done++
				cs[15] = nil
				if done == 16 {
					return
				}
				continue
			}
			out <- v
		}
	}
}
func select8(cs []<-chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var v int
	var ok bool
	done := 0
	for {
		select {
		case v, ok = <-cs[0]:
			if !ok {
				done++
				cs[0] = nil
				if done == 8 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[1]:
			if !ok {
				done++
				cs[1] = nil
				if done == 8 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[2]:
			if !ok {
				done++
				cs[2] = nil
				if done == 8 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[3]:
			if !ok {
				done++
				cs[3] = nil
				if done == 8 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[4]:
			if !ok {
				done++
				cs[4] = nil
				if done == 8 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[5]:
			if !ok {
				done++
				cs[5] = nil
				if done == 8 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[6]:
			if !ok {
				done++
				cs[6] = nil
				if done == 8 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[7]:
			if !ok {
				done++
				cs[7] = nil
				if done == 8 {
					return
				}
				continue
			}
			out <- v
		}
	}
}
func select4(cs []<-chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var v int
	var ok bool
	done := 0
	for {
		select {
		case v, ok = <-cs[0]:
			if !ok {
				done++
				cs[0] = nil
				if done == 4 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[1]:
			if !ok {
				done++
				cs[1] = nil
				if done == 4 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[2]:
			if !ok {
				done++
				cs[2] = nil
				if done == 4 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[3]:
			if !ok {
				done++
				cs[3] = nil
				if done == 4 {
					return
				}
				continue
			}
			out <- v
		}
	}
}
func select2(cs []<-chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var v int
	var ok bool
	done := 0
	for {
		select {
		case v, ok = <-cs[0]:
			if !ok {
				done++
				cs[0] = nil
				if done == 2 {
					return
				}
				continue
			}
			out <- v
		case v, ok = <-cs[1]:
			if !ok {
				done++
				cs[1] = nil
				if done == 2 {
					return
				}
				continue
			}
			out <- v
		}
	}
}
func select1(cs []<-chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var v int
	var ok bool
	done := 0
	for {
		select {
		case v, ok = <-cs[0]:
			if !ok {
				done++
				cs[0] = nil
				if done == 1 {
					return
				}
				continue
			}
			out <- v}
	}
}

func mergeN(chans ...<-chan int) <-chan int {
	r := make(chan int, 1)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(chans))
		for _, c := range chans {
			go func(c <-chan int) {
				for v := range c {
					r <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(r)
	}()

	return r
}

func mergeReflect(chans ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface().(int)
		}
	}()
	return out
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
		}
		close(c)
	}()
	return c
}
func Test_Secltn(t *testing.T) {
	for i := 0; i < 5; i++ {
		chs := make([]<-chan int, 16)
		for i := 0; i < 16; i++ {
			chs[i] = asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
		v := selectn(chs...)
		for rv := range v {
			fmt.Println(rv)
		}
	}
}
func Test_Goroutine(t *testing.T) {
	for i := 0; i < 5; i++ {
		chs := make([]<-chan int, 16)
		for i := 0; i < 16; i++ {
			chs[i] = asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
		v := goroutine(chs...)
		for rv := range v {
			fmt.Println(rv)
		}
	}

}

func BenchmarkMerge(b *testing.B) {
	for _, f := range funcs {
		for n := 1; n <= 4096; n *= 2 {
			chans := make([]<-chan int, n)
			b.Run(fmt.Sprintf("%s/%d", f.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					for i := range chans {
						chans[i] = asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
					}
					b.StartTimer()

					c := f.f(chans...)
					for range c {
					}
				}
			})
		}
	}
}
