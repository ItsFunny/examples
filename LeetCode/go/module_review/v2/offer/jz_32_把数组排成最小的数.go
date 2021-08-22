/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/7/14 8:57 上午
# @File : jz_32_把数组排成最小的数.go
# @Description :
# @Attention :
*/
package offer

import (
	"strconv"
	"strings"
)

type minNums []string

func (m minNums) Len() int {
	return len(m)
}

func (m minNums) Less(i, j int) bool {
	if m[i]+m[j] < m[j]+m[i] {
		return true
	}
	return false
}

func (m minNums) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 关键: 如果 a+b<b+a ,则把 b放前面
func PrintMinNumber(numbers []int) string {
	arrs := make([]string, 0)
	for _, v := range numbers {
		arrs = append(arrs, strconv.Itoa(v))
	}

	// sort.Sort(minNums(arrs))
	PrintMinNumberQSort(arrs, 0, len(arrs)-1)
	r := strings.Builder{}
	for _, v := range arrs {
		r.WriteString(v)
	}
	return r.String()
}

func PrintMinNumberQSort(arrs []string, left, right int) {
	if left < right {
		paration := PrintMinNumberQSortParation(arrs, left, right, func(o1, o2 string) bool {
			return o1+o2 < o2+o1
		})
		PrintMinNumberQSort(arrs, left, paration)
		PrintMinNumberQSort(arrs, paration+1, right)
	}
}

func PrintMinNumberQSortParation(arrs []string, left int, right int, compare func(o1, o2 string) bool) int {
	standard := arrs[left]
	for left < right {
		for ; right > left && arrs[right]+standard >= standard+arrs[right]; right-- {
		}
		arrs[left] = arrs[right]
		for ; left < right && arrs[left]+standard <= standard+arrs[left]; left++ {
		}
		arrs[right] = arrs[left]
	}
	arrs[left] = standard
	return left
}
