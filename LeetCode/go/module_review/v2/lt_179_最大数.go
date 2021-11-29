/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/21 1:35 下午
# @File : lt_179_最大数.go
# @Description :
# @Attention :
*/
package v2

import (
	"sort"
	"strconv"
)

// 关键: 排序: 将2个数间更大的数拍前面,但是需要注意的是 [4,45]: 454 > 445 ,所以排序算法需要定制化
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		x1, x2 := 10, 10
		for x1 <= a {
			x1 *= 10
		}
		for x2 <= b {
			x2 *= 10
		}
		return x1*b+a < x2*a+b
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
